package primes

import (
	"math"
)

func getPrimesBelow(limit int) []int {
	var primes []int // Type definition
	switch {         // As numbers get higher, primes get more sparce. Here we try to initialize a matching slice. Not very exact, but focussed on less growing.
	case limit < 1000:
		primes = make([]int, int(math.Round(limit/2)) // Expanding (or growing) slices is very costly. Let's do it ourselves.
	case limit < 1000000:
		primes = make([]int, limit/4) // Expanding (or growing) slices is very costly. Let's do it ourselves.
	default:
		primes = make([]int, limit/5) // Expanding (or growing) slices is very costly. Let's do it ourselves.
	}
	primes[0] = 2
	primes[1] = 3
	numPrimes := 2
	for i := 5; i <= limit; i += 2 { // We lopen door alle mogelijkheden.
		for n := 1; n <= numPrimes; n++ { // We proberen te delen door eerder gevonden priemgetallen
			if i%primes[n-1] == 0 { // deelbaar!
				break
			}

			if primes[n-1] >= int(math.Round(math.Sqrt(float64(i)))) { // We konden door niets delen, dit is een priemgetal.
				primes[numPrimes] = i
				numPrimes++
				break
			}
		}
	}
	return primes[0:numPrimes]
}
