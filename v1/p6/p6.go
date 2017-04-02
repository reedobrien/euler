// The sum of the squares of the first ten natural numbers is,

// 12 + 22 + ... + 102 = 385 The square of the sum of the first ten natural
// numbers is,

// (1 + 2 + ... + 10)2 = 552 = 3025 Hence the difference between the sum of the
// squares of the first ten natural numbers and the square of the sum is 3025 −
// 385 = 2640.

// Find the difference between the sum of the squares of the first one hundred
// natural numbers and the square of the sum.
package main

import (
	"fmt"
	"math"
	"time"
)

// sqos returns the sum of numbers up to n, squared.
func sqos(n int) int {
	var i, sum int
	for i = 1; i <= n; i++ {
		sum += i
	}
	sq := sum * sum
	return sq
}

// sosq returns the sum of the squares of numbers up to n.
func sosq(n int) int {
	var i, sum int
	for i = 1; i <= n; i++ {
		sum += i * i
	}
	return sum
}

// Difference returns the difference between SquareOfSums and SumOfSquares.
func difference(i int) int {
	return sqos(i) - sosq(i)
}

func main() {
	start := time.Now()
	diff := difference(100):
	fmt.Printf("The difference between sum of sq and sq of sum of the first 100 naturual numbers: %d\n Found in %s.\n", diff, time.Since(start))
}
