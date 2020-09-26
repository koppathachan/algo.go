package memoize

type Function interface {
}

func Fibonacci(n int, cache map[int]int) int {
	if fib, ok := cache[n]; ok {
		return fib
	}
	if n < 2 {
		return n
	}
	fib := Fibonacci(n-1, cache) + Fibonacci(n-2, cache)
	cache[n] = fib
	return fib
}
