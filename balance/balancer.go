package balance

import (
	"sort"
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
	Algorithm     BalancingAlgorithm
	rankFunc      func() *Consumer
}

func NewBalancer(algorithm BalancingAlgorithm) *Balancer {
	b := &Balancer{
		PublisherChan: make(chan interface{}),
		Consumers:     []*Consumer{},
		m:             &sync.RWMutex{},
		load:          0,
		Algorithm:     algorithm,
		rankFunc:      nil,
	}

	//decide the ranking func on init instead of checking on every message
	switch b.Algorithm {
	case RoundRobin:
		b.rankFunc = b.GetRankedRoundRobin
	}
	return b

}

func (p *Balancer) AddConsumer(consumer *Consumer) {
	p.m.Lock()

	for i := int32(0); i < consumer.Weight; i++ {
		p.Consumers = append(p.Consumers, consumer)
	}

	sort.Slice(p.Consumers, func(i, j int) bool {
		return p.Consumers[i].Weight < p.Consumers[j].Weight
	})

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
	return p.rankFunc()
}

func (p *Balancer) GetRankedRoundRobin() *Consumer {
	p.m.RLock()
	defer p.m.RUnlock()
	return p.Consumers[ atomic.LoadUint64(&p.load)%uint64(len(p.Consumers))]
}

func (p *Balancer) Push(data interface{}) {
	atomic.AddUint64(&p.load, 1)
	p.GetRanked().Push(data)
}
