# Tell-all
Tellall (Bellman) is a golang channel library helps you to;

- [x] easily implement pub/sub mechanism
- [x] distribute messages between channels using consistent hash algorithm
- [ ] balance messages between channels using any [balancing algorithms](https://kemptechnologies.com/load-balancer/load-balancing-algorithms-techniques/)
    - [x] round-robin
    - [ ] weighted round robin
    - [ ] fairness load balancing 


### Pub / Sub
First pull the library

```shell script
 go get github.com/tufanbarisyildirim/tellall
```

then, start using like; 

```go
package main
 
 import (
 	"context"
 	"fmt"
 	"github.com/tufanbarisyildirim/tellall/broadcast"
 	"log"
 	"time"
 )
 
 func main() {
 
 	b, err := broadcast.NewPublisher(nil)
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
 
 	go subscriber1.Listen(func(message interface{}, subscriber *broadcast.Subscriber) {
 		fmt.Printf("Sub1[%s] got message  %s from Pub[%s]\n", subscriber.Id, message, subscriber.Publisher.Id)
 	})
 
 	go subscriber2.Listen(func(message interface{}, subscriber *broadcast.Subscriber) {
 		fmt.Printf("Sub2[%s] got message  %s from Pub[%s]\n", subscriber.Id, message, subscriber.Publisher.Id)
 	})
 
 	go func() {
 		for i := 0; i <= 3; i++ {
 			time.Sleep(1 * time.Second)
 			_ = b.Pub(fmt.Sprintf("1: %d", i))
 		}
 	}()
 
 	go func() {
 		for i := 0; i <= 5; i++ {
 			time.Sleep(1 * time.Second)
 			_ = b.Pub(fmt.Sprintf("2: %d", i))
 		}
 		b.Close()
 	}()
 
 	<-b.Done()
 }
```

### Output
```shell script
Sub2[b975a237] got message  1: 0 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  1: 0 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 0 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 0 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  1: 1 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  1: 1 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 1 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 1 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  1: 2 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  1: 2 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 2 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 2 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 3 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 3 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  1: 3 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  1: 3 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 4 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 4 from Pub[183eeb4d23e3e34a]
Sub2[b975a237] got message  2: 5 from Pub[183eeb4d23e3e34a]
Sub1[1b63a6a4] got message  2: 5 from Pub[183eeb4d23e3e34a]
```

## Distributed Hash Table
With distributed consistent hashing you always get same target with same key.

```go
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

````

### Output:
```shell script
got message 0 to ch1
got message 1 to ch1
got message 2 to ch1
got message 3 to ch3
got message 4 to ch2
got message 5 to ch2
got message 6 to ch3
got message 7 to ch1
got message 8 to ch1
got message 9 to ch3
ch1 total message  5
ch2 total message  2
ch3 total message  3
```


## Balancing
Tellall supports channel balancing too

### Round-Robin Balancing
```go
package main

import (
	"fmt"
	"github.com/tufanbarisyildirim/tellall/balance"
)

func main() {

	balancer := balance.NewBalancer()
	consumer1 := balance.NewConsumer(1)
	consumer2 := balance.NewConsumer(1)
	consumer3 := balance.NewConsumer(1)
	balancer.AddConsumer(consumer1)
	balancer.AddConsumer(consumer2)
	balancer.AddConsumer(consumer3)

	consumer1Count := 0
	consumer2Count := 0
	consumer3Count := 0

	go func() {
		for {
			m := <-consumer1.OutChan
			consumer1Count++
			fmt.Println(m)
		}
	}()

	go func() {
		for {
			m := <-consumer2.OutChan
			consumer2Count++
			fmt.Println(m)
		}
	}()

	go func() {
		for {
			m := <-consumer3.OutChan
			consumer3Count++
			fmt.Println(m)
		}
	}()

	for i := 0; i <= 100000; i++ {
		balancer.Push("message")
	}

	fmt.Printf("consumer1 : %d\n", consumer1Count)
	fmt.Printf("consumer2 : %d\n", consumer2Count)
	fmt.Printf("consumer3 : %d\n", consumer3Count)

}
```



### Output
```shell script
consumer1 : 33333
consumer2 : 33334
consumer3 : 33334
```