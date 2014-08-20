package main

import (
	"errors"
	"fmt"
)

func makeTriple(m, n int) (a, b, c int, err error) {
	if m > n {
		err := errors.New("m must be less than n")
		return 0, 0, 0, err
	}
	a = n*n - m*m
	b = 2 * n * m
	c = n*n + m*m
	return a, b, c, nil
}

func main() {
	done := false
	m, n := 1, 2
	for done == false {
		a, b, c, err := makeTriple(m, n)
		if err != nil {
			panic(err)
		}
		if a+b+c == 1000 {
			fmt.Println(a * b * c)
			done = true
		} else {
			n = n + 1
		}
		if a+b+c > 1000 {
			m = m + 1
			n = m + 2
		}
	}
}
