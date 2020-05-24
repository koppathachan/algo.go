package main

import (
	"fmt"

	"github.com/sasidakh/algo.go/eratosthenes"
)

func main() {
	var n = 5
	fmt.Println(eratosthenes.Seive(n))
}
