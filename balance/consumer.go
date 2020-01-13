package balance

import (
	"github.com/tufanbarisyildirim/tellall/utils"
	"sync/atomic"
)

type Consumer struct {
	Id      string
	OutChan chan interface{}
	Weight  int32
	load    uint64
}

func NewConsumer(weight int32) *Consumer {
	return &Consumer{
		Id:      utils.RandHexId(4),
		OutChan: make(chan interface{}),
		Weight:  weight,
		load:    0,
	}
}

func (c *Consumer) Push(data interface{}) {
	atomic.AddUint64(&c.load, 1)
	c.OutChan <- data
}

func (c *Consumer) GetLoad() uint64 {
	return atomic.LoadUint64(&c.load)
}
