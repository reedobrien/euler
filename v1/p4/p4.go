//A palindromic number reads the same both ways. The largest palindrome made
//from the product of two 2-digit numbers is 9009 = 91 × 99.
//
//Find the largest palindrome made from the product of two 3-digit numbers.
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
	"time"
)

var a, b, aMin, bMin, largest, product int

func strReverse(s string) string {
	runes := []rune(s)
	var rev []rune
	for i := len(s) - 1; i >= 0; i-- {
		rev = append(rev, runes[i])
	}
	return string(rev)
}

func sortReverse(s string) string {
	digits := strings.Split(s, "")
	sort.Reverse(sort.StringSlice(digits))
	return strings.Join(digits, "")
}

func main() {
	aMin = 99
	bMin = 99
	start := time.Now()
	for a = 999; a > aMin; a-- {
		for b = a; b > bMin; b-- {
			product = a * b
			pStr := strconv.Itoa(product)
			if pStr == strReverse(pStr) {
				if product > largest {
					largest = product
					break
				}
			}
		}
	}
	fmt.Printf("Largest palindrome product of two three digit numbers: %d\n Found in: %s seconds\n", largest, time.Since(start))
}
