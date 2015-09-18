// The prime factors of 13195 are 5, 7, 13 and 29.

// What is the largest prime factor of the number 600,851,475,143 ?
package main

import (
	"fmt"
	"time"
)

var (
	large, d int64
	number   int64 = 600851475143
)

func main() {
	start := time.Now()
	factors := primeFactors(number)
	elapsed := time.Since(start)
	// last one is the largest
	large = factors[len(factors)-1]
	fmt.Printf("Largest prime factor: %d\nFound in %s\n", large, elapsed)
}

// Dumb factoring and prime testing
func primeFactors(n int64) []int64 {
	var factors []int64
	d = 2
	for n > 1 {
		for n%d == 0 {
			factors = append(factors, d)
			n /= d
		}
		d += 1
		if d*d > n {
			if n > 1 {
				factors = append(factors, n)
			}
			return factors
		}
	}
	return factors
}
