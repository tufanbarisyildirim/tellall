# bellman
Bellman is a golang channel library broadcasts any message to multiple channels/subscribers. 


### Example Usage
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

	b := bellman.NewBellman(context.Background())

	subscriber1, err := b.NewSub(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	subscriber2, err := b.NewSub(context.Background())
	if err != nil {
		log.Fatal(err)
	}

	go subscriber1.Listen(func(message interface{}) {
		fmt.Printf("Sub1[%s] %s\n", subscriber1.Id, message)
	})

	go subscriber2.Listen(func(message interface{}) {
		fmt.Printf("Sub2[%s] %s\n", subscriber2.Id, message)
	})

	go func() {
		for i := 0; i <= 3; i++ {
			time.Sleep(1 * time.Second)
			_ = b.Pub(fmt.Sprintf("Publisher 1: %d", i))
		}
	}()

	go func() {
		for i := 0; i <= 5; i++ {
			time.Sleep(1 * time.Second)
			_ = b.Pub(fmt.Sprintf("Publisher 2: %d", i))
		}
		b.Close()
	}()

	<-b.Done()
}
```

### Output
```shell script
Sub2[83614df7] Publisher 2: 0
Sub1[c22496d6] Publisher 2: 0
Sub1[c22496d6] Publisher 1: 0
Sub2[83614df7] Publisher 1: 0
Sub2[83614df7] Publisher 1: 1
Sub2[83614df7] Publisher 2: 1
Sub1[c22496d6] Publisher 1: 1
Sub1[c22496d6] Publisher 2: 1
Sub2[83614df7] Publisher 1: 2
Sub1[c22496d6] Publisher 1: 2
Sub2[83614df7] Publisher 2: 2
Sub1[c22496d6] Publisher 2: 2
Sub1[c22496d6] Publisher 1: 3
Sub2[83614df7] Publisher 1: 3
Sub2[83614df7] Publisher 2: 3
Sub1[c22496d6] Publisher 2: 3
Sub2[83614df7] Publisher 2: 4
Sub1[c22496d6] Publisher 2: 4
Sub2[83614df7] Publisher 2: 5
Sub1[c22496d6] Publisher 2: 5
```