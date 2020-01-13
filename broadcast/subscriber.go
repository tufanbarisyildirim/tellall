package broadcast

import (
	"context"
	"crypto/rand"
	"fmt"
)

type Subscriber struct {
	ctx       context.Context
	ch        chan interface{}
	closed    chan struct{}
	Id        string
	Publisher *Publisher
}

func NewSubscriber(ctx context.Context, publisher *Publisher) (*Subscriber, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	b := make([]byte, 4)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("error generating subscriber id :%s", err)
	}
	return &Subscriber{
		ctx:       ctx,
		ch:        make(chan interface{}),
		closed:    make(chan struct{}),
		Id:        fmt.Sprintf("%x", b),
		Publisher: publisher,
	}, nil
}

func (s *Subscriber) Unsub(bellman *Publisher) bool {
	return bellman.Kick(s)
}

func (s *Subscriber) Sub(bellman *Publisher) {
	bellman.Sub(s)
}

func (s *Subscriber) Listen(listener func(message interface{}, subscriber *Subscriber)) {
outer:
	for {
		select {
		case <-s.ctx.Done():
			break outer
		case <-s.closed:
			break outer
		case m := <-s.ch:
			listener(m, s)
			break
		}
	}
}
