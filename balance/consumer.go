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

func NewConsumer() (*Consumer, error) {
	id, err := utils.RandHexId(4)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		Id:      id,
		OutChan: make(chan interface{}),
		Weight:  1,
		load:    0,
	}, nil
}

func NewWeightedConsumer(weight int32) (*Consumer, error) {
	c, err := NewConsumer()
	if err != nil {
		return nil, err
	}
	c.Weight = weight
	return c, nil
}

func (c *Consumer) Push(data interface{}) {
	atomic.AddUint64(&c.load, 1)
	c.OutChan <- data
}

func (c *Consumer) GetLoad() uint64 {
	return atomic.LoadUint64(&c.load)
}
