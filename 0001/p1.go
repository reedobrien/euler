package main

import "fmt"

var (
	max = 1000
	Σ   = 0
)

func main() {
	for i := 3; i < max; i++ {
		if i%3 == 0 || i%5 == 0 {
			Σ += i
		}
	}
	fmt.Printf("Sum of multiples of 3 or 5 below %d is %d\n", max, Σ)
}
