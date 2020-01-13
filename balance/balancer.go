package balance

import (
	"sync"
	"sync/atomic"
)

type BalancingAlgorithm uint8

const (
	RoundRobin = BalancingAlgorithm(iota)
)

type Balancer struct {
	PublisherChan chan interface{}
	Consumers     []*Consumer
	m             *sync.RWMutex
	load          uint64
}

func NewBalancer() *Balancer {
	return &Balancer{
		PublisherChan: make(chan interface{}),
		Consumers:     []*Consumer{},
		m:             &sync.RWMutex{},
		load:          0,
	}
}

func (p *Balancer) AddConsumer(consumer *Consumer) {
	p.m.Lock()
	p.Consumers = append(p.Consumers, consumer)
	p.m.Unlock()
}

func (p *Balancer) RemoveConsumer(consumer *Consumer) {
	p.m.Lock()
	i := 0
	for _, c := range p.Consumers {
		if c != consumer {
			p.Consumers[i] = c
			i++
		}
	}
	p.Consumers = p.Consumers[:i]
	p.m.Unlock()
}

func (p *Balancer) GetRanked() *Consumer {
	p.m.RLock()
	defer p.m.RUnlock()
	return p.Consumers[ atomic.LoadUint64(&p.load)%uint64(len(p.Consumers))]
}

func (p *Balancer) Push(data interface{}) {
	atomic.AddUint64(&p.load, 1)
	p.GetRanked().Push(data)
}
