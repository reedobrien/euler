package utils

import "math/big"

var a, b *big.Int

func Fibgen() chan string {
	c := make(chan string)
	go func() {
		for a, b := big.NewInt(0), big.NewInt(1); ; a, b = b, a.Add(a, b) {
			c <- a.String()
		}
	}()
	return c
}
