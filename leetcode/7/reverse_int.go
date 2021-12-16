package main

import (
	"math"
)

func reverse(x int) int {
	res := 0
	sign := 1
	if x < 0 {
		sign = -1
	}

	for x != 0 {
		pop := x % 10
		x = x / 10

		res = res*10 + pop
	}

	if res > 0 && sign < 0 {
		return res * sign
	}
	if res < math.MinInt32 || res > math.MaxInt32 {
		return 0
	}
	return res
}
