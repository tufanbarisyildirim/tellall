package balance

import (
	"sync/atomic"
	"time"

	"github.com/tufanbarisyildirim/balancer"
	"github.com/tufanbarisyildirim/tellall/utils"
)

//Consumer data consumer
type Consumer struct {
	balancer.Node
	ID       string
	OutChan  chan interface{}
	Weight   int32
	load     int64
	consumed uint64
}

//NewConsumer create a new consumer with defaults
func NewConsumer() (*Consumer, error) {
	id, err := utils.RandHexId(4)

	if err != nil {
		return nil, err
	}

	return &Consumer{
		ID:      id,
		OutChan: make(chan interface{}),
		Weight:  1,
		load:    0,
	}, nil
}

//NewWeightedConsumer create a consumer with a predefined weight
func NewWeightedConsumer(weight int32) (*Consumer, error) {
	c, err := NewConsumer()
	if err != nil {
		return nil, err
	}
	c.Weight = weight
	return c, nil
}

//Push push data to consumer
func (c *Consumer) Push(data interface{}) {
	atomic.AddInt64(&c.load, 1)
	c.OutChan <- data
	//TODO(tufan): find a way to have a callback from worker to calculate time
}

//IsHealthy check if consumer still up
func (c *Consumer) IsHealthy() bool {
	return true // check if channel stil alive?
}

//TotalRequest total consumed data count
func (c *Consumer) TotalRequest() uint64 {
	return c.consumed
}

//AverageResponseTime average time for one job
func (c *Consumer) AverageResponseTime() time.Duration {
	return time.Second * 1 //return same for all for now
}

//Load get current load of
func (c *Consumer) Load() int64 {
	return atomic.LoadInt64(&c.load)
}

//NodeID return consumer ID for balancer
func (c *Consumer) NodeID() string {
	return c.ID
}
