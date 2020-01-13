package bellman

import (
	"context"
	"crypto/rand"
	"fmt"
	"sync"
)

type Publisher struct {
	eventObserver chan Event
	ctx           context.Context
	closed        chan struct{}
	Id            string
}

func NewPublisher(ctx context.Context) (*Publisher, error) {
	if ctx == nil {
		ctx = context.Background()
	}

	randBytes := make([]byte, 8)
	_, err := rand.Read(randBytes)
	if err != nil {
		return nil, fmt.Errorf("error generating subscriber id :%s", err)
	}

	p := &Publisher{
		ctx:           ctx,
		eventObserver: make(chan Event),
		closed:        make(chan struct{}),
		Id:            fmt.Sprintf("%x", randBytes),
	}

	go func() {
		var m sync.RWMutex
		subscribers := map[string]*Subscriber{}
	outer:
		for {
			select {
			case <-ctx.Done():
				break outer
			case e := <-p.eventObserver:
				switch e.Signal {
				case SIGPUB:
					go func() {
						m.Lock()
					subsLoop:
						for _, subscriber := range subscribers {
							select {
							case <-subscriber.ctx.Done():
								subscriber.Unsub(p)
								continue subsLoop
							case subscriber.ch <- e.Payload:
								break
							case <-subscriber.closed:
								continue subsLoop
							}
						}
						m.Unlock()
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

		close(p.closed)

	}()

	return p,nil
}

func (p *Publisher) Pub(message interface{}) error {
	p.eventObserver <- Event{
		Signal:  SIGPUB,
		Payload: message,
	}

	return nil
}

func (p *Publisher) NewSub(ctx context.Context) (*Subscriber, error) {
	s, err := NewSubscriber(ctx, p)
	if err != nil {
		return nil, err
	}
	p.Sub(s)
	return s, nil
}

func (p *Publisher) Sub(subscriber *Subscriber) {
	p.eventObserver <- Event{
		Signal:  SIGSUB,
		Payload: subscriber,
	}
}

func (p *Publisher) Kick(subscriber *Subscriber) bool {
	return p.KickId(subscriber.Id)
}

func (p *Publisher) KickId(subscriberId string) bool {
	p.eventObserver <- Event{
		Signal:  SIGUNSUB,
		Payload: subscriberId,
	}
	return true
}

func (p *Publisher) Close() bool {
	p.eventObserver <- Event{
		Signal:  SIGTERM,
		Payload: nil,
	}
	return true
}

func (p *Publisher) Done() <-chan struct{} {
	return p.closed
}
