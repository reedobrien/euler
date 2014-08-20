/* euler 10 Seive of Eratosthenes
http://en.wikipedia.org/wiki/Sieve_of_Eratosthenes#Implementation
*/
package main

import "fmt"

const max int64 = 2000000
const expected int64 = 142913828922

var Σ int64 = 0
var A = make([]bool, max+1, max+1)

func main() {
	A[0], A[1] = true, true

	for i := 0; i < len(A); i++ {
		// start at next false
		if !A[i] {
			// sum if not compound
			Σ += int64(i)

			for y := i + i; y < len(A); y += i {
				// multiples are compound
				A[y] = true
			}
		}
	}
	fmt.Printf("sum of primes up to %d is %d and %v\n", max, Σ, Σ == expected)

}
