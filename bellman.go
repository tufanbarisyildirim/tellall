package bellman

import (
	"context"
	"sync"
)

type Bellman struct {
	eventObserver chan Event
	ctx           context.Context
	w             sync.WaitGroup
}

func NewBellman(ctx context.Context) *Bellman {
	if ctx == nil {
		ctx = context.Background()
	}

	b := &Bellman{
		ctx:           ctx,
		eventObserver: make(chan Event),
	}
	b.w.Add(1)

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
								continue subsLoop
							case subscriber.ch <- e.Payload:
							}

						}
						m.RUnlock()
					}()
					break
				case SIGSUB:
					m.Lock()
					b.w.Add(1)
					subscribers[e.Payload.(*Subscriber).Id] = e.Payload.(*Subscriber)
					m.Unlock()
					break
				case SIGUNSUB:
					m.Lock()
					if s, ok := subscribers[e.Payload.(string)]; ok {
						b.w.Done()
						close(s.ch)
						delete(subscribers, e.Payload.(string))
					}
					m.Unlock()
					break
				case SIGTERM:
					b.w.Done()
					break outer
				}
			}
		}

		for _, subscriber := range subscribers {
			b.w.Done()
			close(subscriber.ch)
		}

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

func (b *Bellman) Sub(ctx context.Context) (*Subscriber, error) {
	s, err := NewSubscriber(ctx)
	if err != nil {
		return nil, err
	}
	b.eventObserver <- Event{
		Signal:  SIGSUB,
		Payload: s,
	}
	return s, nil
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

func (b *Bellman) Wait() {
	b.w.Wait()
}
