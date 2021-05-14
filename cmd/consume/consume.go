package main

import (
	"fmt"

	"github.com/sasidakh/algo.go/pubsub"
)

func main() {
	c := pubsub.Subscribe("q1")
	for {
		select {
		case x := <-c:
			fmt.Println(x)
		}
	}
}
