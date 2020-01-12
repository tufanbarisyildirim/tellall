package bellman

import (
	"context"
	"crypto/rand"
	"fmt"
)

type Subscriber struct {
	ctx context.Context
	ch  chan interface{}
	Id  string
}

func NewSubscriber(ctx context.Context) (*Subscriber, error) {

	if ctx == nil {
		ctx = context.Background()
	}

	b := make([]byte, 16)
	_, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("error generating subscriber id :%s", err)
	}
	return &Subscriber{
		ctx: ctx,
		ch:  make(chan interface{}),
		Id:  fmt.Sprintf("%x", b),
	}, nil
}

func (s *Subscriber) Unsub(bellman *Bellman) bool {
	return bellman.UnSub(s)
}

func (s *Subscriber) Listen(listener func(message interface{})) {
outer:
	for {
		select {
		case <-s.ctx.Done():
			break outer
		case m := <-s.ch:
			listener(m)

			break
		}
	}
}
