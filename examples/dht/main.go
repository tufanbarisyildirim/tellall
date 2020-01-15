package main

import (
	"fmt"
	"github.com/tufanbarisyildirim/tellall/dht"
)

func main() {

	t := dht.NewHashTable()
	ch1 := make(chan interface{})
	ch2 := make(chan interface{})
	ch3 := make(chan interface{})

	t.Add(ch1)
	t.Add(ch2)
	t.Add(ch3)
	ch1Count := 0
	ch2Count := 0
	ch3Count := 0
	go func() {
		for {
			select {
			case i := <-ch1:
				ch1Count++
				fmt.Printf(" got message %d to ch1\n", i)
				break
			case i := <-ch2:
				ch2Count++
				fmt.Printf(" got message %d to ch2\n", i)
				break
			case i := <-ch3:
				ch3Count++
				fmt.Printf(" got message %d to ch3\n", i)
				break
			}
		}
	}()

	for i := 0; i < 10; i++ {
		t.Push(uint64(i), i)
	}

	fmt.Println("ch1 total message ", ch1Count)
	fmt.Println("ch2 total message ", ch2Count)
	fmt.Println("ch3 total message ", ch3Count)
}
