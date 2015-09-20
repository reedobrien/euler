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

func sqos(n int64) int64 {
	var i, sum int64
	for i = 1; i <= n; i++ {
		sum += i
	}
	sq := math.Pow(float64(sum), 2)
	return int64(sq)
}

func sosq(n int64) int64 {
	var i, sum int64
	for i = 1; i <= n; i++ {
		sum += int64(math.Pow(float64(i), 2))
	}
	return sum
}

func main() {
	start := time.Now()
	diff := sqos(100) - sosq(100)
	fmt.Printf("The difference between sum of sq and sq of sum of the first 100 naturual numbers: %d\n Found in %s.\n", diff, time.Since(start))
}
