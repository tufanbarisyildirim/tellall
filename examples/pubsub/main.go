package main

import (
	"context"
	"fmt"
	"github.com/tufanbarisyildirim/tellall/pubsub"
	"log"
	"time"
)

func main() {

	b, err := pubsub.NewPublisher(nil)
	if err != nil {
		log.Fatal(err)
	}

	subscriber1, err := b.NewSub(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	subscriber2, err := b.NewSub(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go subscriber1.Listen(func(message interface{}, subscriber *pubsub.Subscriber) {
		fmt.Printf("Sub1[%s] got message  %s from Pub[%s]\n", subscriber.Id, message, subscriber.Publisher.Id)
	})

	go subscriber2.Listen(func(message interface{}, subscriber *pubsub.Subscriber) {
		fmt.Printf("Sub2[%s] got message  %s from Pub[%s]\n", subscriber.Id, message, subscriber.Publisher.Id)
	})

	go func() {
		for i := 0; i <= 3; i++ {
			time.Sleep(300 * time.Millisecond)
			_ = b.Pub(fmt.Sprintf("1: %d", i))
		}
	}()

	go func() {
		for i := 0; i <= 5; i++ {
			time.Sleep(250 * time.Millisecond)
			_ = b.Pub(fmt.Sprintf("2: %d", i))
		}
		b.Close()
	}()

	<-b.Done()
}
