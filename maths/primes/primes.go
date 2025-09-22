package primes

import (
	"math"
)

type NumbersMap map[int]bool

func (s NumbersMap) add(num int) {
	s[num] = true
}
func (s NumbersMap) del(num int) {
	delete(s, num)
}

func getPrimesBelow(limit int) []int {
	var primes = make([]int, limit/2) // No endlessly growing slices, expanding is costly
	primes[0] = 2                     // a given
	primes[1] = 3                     // a given
	numPrimes := 2
	for i := 5; i <= limit; i += 2 { // Traverse all numbers that could be prime
		for n := 1; n <= numPrimes; n++ {
			if i%primes[n-1] == 0 { // Divisiable, so this number (i) is not prime!
				break
			}

			if primes[n-1] >= int(math.Round(math.Sqrt(float64(i)))) { // End of the possible divisors, this is a prime!
				primes[numPrimes] = i
				numPrimes++
				break
			}
		}

	}
	return primes[0:numPrimes]
}
