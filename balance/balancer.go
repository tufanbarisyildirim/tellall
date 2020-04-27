package balance

import (
	"github.com/tufanbarisyildirim/balancer"
)

//Balancer channel balancer
type Balancer struct {
	PublisherChan chan interface{}
	b             *balancer.Balancer
}

//NewBalancer create new balancer
func NewBalancer(selectionPolicy balancer.SelectionPolicy) *Balancer {
	b := &Balancer{
		PublisherChan: make(chan interface{}),
		b:             balancer.NewBalancer(),
	}
	return b
}

//AddConsumer add a new channel consumer
func (b *Balancer) AddConsumer(consumer *Consumer) {
	b.b.Add(consumer)
}

//RemoveConsumer a consumer from nodes
func (b *Balancer) RemoveConsumer(consumer *Consumer) {
	b.b.Remove(consumer.NodeID())
}

//Push push data!
func (b *Balancer) Push(data interface{}) {
	b.b.Next("0").(*Consumer).Push(data)
}
