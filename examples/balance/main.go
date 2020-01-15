package main

import (
	"fmt"
	"github.com/tufanbarisyildirim/tellall/balance"
)

func main() {

	balancer := balance.NewBalancer(balance.RoundRobin)
	consumer1, _ := balance.NewConsumer()
	consumer2, _ := balance.NewWeightedConsumer(2)
	consumer3, _ := balance.NewWeightedConsumer(99)
	balancer.AddConsumer(consumer1)
	balancer.AddConsumer(consumer2)
	balancer.AddConsumer(consumer3)

	consumer1Count := 0
	consumer2Count := 0
	consumer3Count := 0

	go func() {
		for {
			<-consumer1.OutChan
			consumer1Count++
		}
	}()

	go func() {
		for {
			<-consumer2.OutChan
			consumer2Count++
		}
	}()

	go func() {
		for {
			<-consumer3.OutChan
			consumer3Count++
		}
	}()

	for i := 0; i <= 100000; i++ {
		balancer.Push("message")
	}

	fmt.Printf("consumer1 : %d\n", consumer1Count)
	fmt.Printf("consumer2 : %d\n", consumer2Count)
	fmt.Printf("consumer3 : %d\n", consumer3Count)

}
