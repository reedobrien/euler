// By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see
// that the 6th prime is 13.

// What is the 10 001st prime number?
package main

import (
	"fmt"
	"math/big"
	"os"
	"time"
)

var (
	nth, i int64
)

func main() {
	start := time.Now()
	nth = 4
	for i = 7; ; i += 2 {
		if big.NewInt(i).ProbablyPrime(20) {
			nth++
			if nth == 10001 {
				fmt.Printf("The 10,001st prime is: %d\nFound in %s\n", i, time.Since(start))
				os.Exit(0)
			}
		}
	}

}
