package bellman

import (
	"context"
	"sync"
)

type Bellman struct {
	eventObserver chan Event
	ctx           context.Context
	closed        chan struct{}
}

func NewBellman(ctx context.Context) *Bellman {
	if ctx == nil {
		ctx = context.Background()
	}

	b := &Bellman{
		ctx:           ctx,
		eventObserver: make(chan Event),
		closed:        make(chan struct{}),
	}

	go func() {
		var m sync.RWMutex
		subscribers := map[string]*Subscriber{}
	outer:
		for {
			select {
			case <-ctx.Done():
				break outer
			case e := <-b.eventObserver:
				switch e.Signal {
				case SIGPUB:
					go func() {
						m.RLock()
					subsLoop:
						for _, subscriber := range subscribers {
							select {
							case <-subscriber.ctx.Done():
								subscriber.Unsub(b)
								continue subsLoop
							case subscriber.ch <- e.Payload:
								break
							case <-subscriber.closed:
								continue subsLoop
							}
						}
						m.RUnlock()
					}()
					break
				case SIGSUB:
					m.Lock()
					subscribers[e.Payload.(*Subscriber).Id] = e.Payload.(*Subscriber)
					m.Unlock()
					break
				case SIGUNSUB:
					m.Lock()
					if s, ok := subscribers[e.Payload.(string)]; ok {
						close(s.ch)
						close(s.closed)
						delete(subscribers, e.Payload.(string))
					}
					m.Unlock()
					break
				case SIGTERM:
					break outer
				}
			}
		}

		for _, subscriber := range subscribers {
			close(subscriber.closed)
			close(subscriber.ch)
		}

		close(b.closed)

	}()

	return b
}

func (b *Bellman) Pub(message interface{}) error {
	b.eventObserver <- Event{
		Signal:  SIGPUB,
		Payload: message,
	}

	return nil
}

func (b *Bellman) NewSub(ctx context.Context) (*Subscriber, error) {
	s, err := NewSubscriber(ctx)
	if err != nil {
		return nil, err
	}
	b.Sub(s)
	return s, nil
}

func (b *Bellman) Sub(subscriber *Subscriber) {
	b.eventObserver <- Event{
		Signal:  SIGSUB,
		Payload: subscriber,
	}
}

func (b *Bellman) UnSub(subscriber *Subscriber) bool {
	return b.UnSubId(subscriber.Id)
}

func (b *Bellman) UnSubId(subscriberId string) bool {
	b.eventObserver <- Event{
		Signal:  SIGUNSUB,
		Payload: subscriberId,
	}
	return true
}

func (b *Bellman) Close() bool {
	b.eventObserver <- Event{
		Signal:  SIGTERM,
		Payload: nil,
	}
	return true
}

func (b *Bellman) Done() <-chan struct{} {
	return b.closed
}
