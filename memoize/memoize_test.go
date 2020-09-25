package memoize_test

import (
	"github.com/sasidakh/algo.go/memoize"
	"testing"
)

func Fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return Fibonacci(n-1) + Fibonacci(n-2)
}

func BenchmarkFibonacci(b *testing.B) {
	cache := make(map[int]int)
	for i := 0; i < b.N; i++ {
		memoize.Fibonacci(40, cache)
	}
}

func BenchmarkSlowFibonacci(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Fibonacci(40)
	}
}
