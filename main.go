package main

import (
	"fmt"
	"github.com/sasidakh/algo.go/memoize"
	"time"
)

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func main() {
	start := time.Now()
	res := Fibonacci(40)
	elapsed := time.Since(start)

	cache := make(map[int]int)

	start2 := time.Now()
	res2 := memoize.Fibonacci(40, cache)
	elapsed2 := time.Since(start2)

	start3 := time.Now()
	res3 := memoize.Fibonacci(40, cache)
	elapsed3 := time.Since(start3)

	fmt.Println("Time : ", elapsed)
	fmt.Println("Result: ", res)

	fmt.Println("Time : ", elapsed2)
	fmt.Println("Result: ", res2)

	fmt.Println("Time : ", elapsed3)
	fmt.Println("Result: ", res3)
}
