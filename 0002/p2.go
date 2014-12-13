package main

import (
	"fmt"
	"log"
	"strconv"

	"github.com/reedobrien/euler/utils"
)

var (
	max int64 = 4000000
	Σ   int64 = 0
	i   int64 = 0
	s         = ""
	err error
)

func main() {
	fib := utils.Fibgen()
	for i <= max {
		s = <-fib
		i, err = strconv.ParseInt(s, 10, 32)
		if err != nil {
			log.Fatal(err)
		}
		if i%2 == 0 {
			Σ += i
		}
	}
	fmt.Printf("Answer: %d\n", Σ)
}
