# bellman
Bellman is a golang channel library broadcasts any message to multiple channels/subscribers. 


### Example Usage

First pull the library

```shell script
 go get github.com/tufanbarisyildirim/bellman
```

then, start using like; 

```go
package main

import (
	"context"
	"fmt"
	"github.com/tufanbarisyildirim/bellman"
	"log"
	"time"
)

func main() {

	b, err := bellman.NewPublisher(nil)
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

	go subscriber1.Listen(func(message interface{}, subscriber *bellman.Subscriber) {
		fmt.Printf("Sub1[%s] got message  %s from Pub[%s]\n", subscriber.Id, message, subscriber.Publisher.Id)
	})

	go subscriber2.Listen(func(message interface{}, subscriber *bellman.Subscriber) {
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