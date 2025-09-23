package primes

import (
	"math"
)

// getPrimeFactors finds the primefactors of a given number and returns them in a slice
func getPrimeFactors(number int) []int {
	// this function will get the prime	factors.
	square := int(math.Round(math.Sqrt(float64(number))))
	primes := getPrimesBelow(square)
	var primefactors []int
	for _, prime := range primes {
		for {
			if number == 0 {
				return primefactors
			}
			if number%prime == 0 {
				number = number / prime
				primefactors = append(primefactors, prime)
			} else {
				break
			}
		}
	}
	return primefactors
}

// getPrimesBelow calculates all the primes below the given limt and returns them in a slice
func getPrimesBelow(limit int) []int {
	var primes []int
	if limit < 1000 {
		primes = make([]int, limit/2) // Slices expanden is erg duur, zelf doen in dit geval dus!
	} else {
		primes = make([]int, limit/4) // Slices expanden is erg duur, zelf doen in dit geval dus!
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
