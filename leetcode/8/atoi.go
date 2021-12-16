package myatoi

import (
	"math"
	"strconv"
)

var (
	s0     = '0'
	s9     = '9'
	sminus = '-'
	splus  = '+'
	ws     = ' '
)

func myAtoi(s string) int {
	res := make([]rune, 0, 10)
	for _, symbol := range s {
		if symbol == ws {
			if len(res) > 0 {
				break
			}
			continue
		}
		if symbol >= s0 && symbol <= s9 {
			res = append(res, symbol)
			continue
		}
		if symbol == sminus || symbol == splus {
			if len(res) > 0 {
				break
			}
			res = append(res, symbol)
			continue
		}
		if len(res) == 0 {
			return 0
		}
		if len(res) > 0 {
			break
		}
	}
	r, _ := strconv.Atoi(string(res))

	if r < math.MinInt32 {
		return math.MinInt32
	}
	if r > math.MaxInt32 {
		return math.MaxInt32
	}
	return r
}
