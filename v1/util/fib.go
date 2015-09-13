package util

import "math/big"

var a, b *big.Int

// Fibgen returns an unbuffered channel that generates the next number in the
// fibonacci sequence each time it is received from.
func Fibgen() chan string {
	c := make(chan string)
	go func() {
		for a, b := big.NewInt(0), big.NewInt(1); ; a, b = b, a.Add(a, b) {
			c <- a.String()
		}
	}()
	return c
}
