package eratosthenes

func findNextPrime(seive map[int]bool, from int, max int) int {
	i := from
	for {
		if _, ok := seive[i]; !ok {
			return i
		}
		i++
	}
}

func newSeive(n int) map[int]bool {
	seive := make(map[int]bool)
	for i := 2; i <= n; i++ {
		seive[i] = true
	}
	return seive
}

func deleteMultipleOf(p int, n int, seive map[int]bool) {
	i := 0
	for key := p * p; i < n && key <= n; key = p * (p + i) {
		if seive[key] {
			delete(seive, key)
		}
		i++
	}

}

func Seive(n int) []int {
	seive := newSeive(n)
	p := 2
	for p <= n {
		deleteMultipleOf(p, n, seive)
		p = findNextPrime(seive, p+1, n)
	}
	primes := make([]int, len(seive))

	i := 0
	for key := range seive {
		primes[i] = key
		i++
	}
	return primes
}
