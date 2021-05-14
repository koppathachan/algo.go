package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/sasidakh/algo.go/pubsub"
)

func main() {
	noquit := make(chan int)
	c := pubsub.Subscribe("q1")
	go func() {
		for {
			select {
			case x := <-c:
				fmt.Println("got ", x)
			}
		}
	}()
	go func() {
		for {
			pub := rand.Intn(10)
			time.Sleep(time.Second)
			fmt.Println("publishing ", pub)
			pubsub.Publish(pub, "q1")
		}
	}()
	<-noquit
}
