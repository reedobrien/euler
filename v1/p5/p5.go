// 2520 is the smallest number that can be divided by each of the numbers from
// 1 to 10 without any remainder.

// What is the smallest positive number that is evenly divisible by all of the
// numbers from 1 to 20?
package main

import (
	"fmt"
	"os"
	"time"
)

var (
	n int
)

func main() {
	start := time.Now()
	for n = 20; ; n += 2 {
		if divisibleByAll(n) == true {
			fmt.Printf("Smallest positive number divisible by all numbers 1-20: %d\nFound in %s\n", n, time.Since(start))
			os.Exit(0)
		}
	}
}

func divisibleByAll(n int) bool {
	for i := 20; i > 0; i-- {
		if n%i != 0 {
			return false
		}
	}
	return true
}
