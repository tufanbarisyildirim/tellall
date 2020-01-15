package broadcast

import (
	"context"
	"github.com/tufanbarisyildirim/tellall/utils"
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

	id, err := utils.RandHexId(4)
	if err != nil {
		return nil, err
	}

	return &Subscriber{
		ctx:       ctx,
		ch:        make(chan interface{}),
		closed:    make(chan struct{}),
		Id:        id,
		Publisher: publisher,
	}, nil
}

func (s *Subscriber) Unsub(publisher *Publisher) bool {
	return publisher.Kick(s)
}

func (s *Subscriber) Sub(publisher *Publisher) {
	publisher.Sub(s)
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
