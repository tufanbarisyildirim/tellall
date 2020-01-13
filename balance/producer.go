package balance

import "sync"

type Producer struct {
	PublisherChan chan interface{}
	Consumers     []*Consumer
	m             *sync.RWMutex
}

func (p *Producer) AddConsumer(consumer *Consumer) {
	p.m.Lock()
	p.Consumers = append(p.Consumers, consumer)
	p.m.Unlock()
}

func (p *Producer) RemoveConsumer(consumer Consumer) {
	p.m.Lock()
	i := 0
	for _, c := range h.Bucket {
		if c != ch {
			p.Bucket[i] = c
			i++
		}
	}
	h.Bucket = h.Bucket[:i]
	p.m.Unlock()
}