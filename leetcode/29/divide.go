package main

import (
	"fmt"
	"math"
)

func divide(dividend int, divisor int) int {

	minus := (dividend < 0) != (divisor < 0)

	res := 0
	divisor = abs(divisor)
	dividend = abs(dividend)
	if dividend == math.MaxInt32+1 && !minus {
		dividend--
	}
	if divisor == math.MaxInt32+1 && !minus {
		divisor--
	}

	for dividend >= divisor {
		dividend -= divisor
		res++
	}

	if minus {
		res = -res
	} else {

	}
	return res
}

func abs(m int) int {
	if m < 0 {
		return -m
	}
	return m
}

func main() {
	fmt.Println(divide(1024001314, 1515095570))
}
