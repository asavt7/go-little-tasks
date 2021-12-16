package myatoi

import (
	"testing"
)

func TestName(t *testing.T) {

	tt := []struct {
		in  string
		res int
	}{{
		in:  "   -42",
		res: -42,
	}, {
		in:  "4193 with words",
		res: 4193,
	}, {
		in:  " with words 4193",
		res: 0,
	}, {
		in:  "-91283472332",
		res: -2147483648,
	}, {
		in:  "3.14159",
		res: 3,
	},
		{
			in:  "   +0 123",
			res: 0,
		}, {
			in:  "-5-",
			res: -5,
		}, {
			in:  " b11228552307",
			res: 0,
		}}

	for _, s := range tt {
		t.Run(s.in, func(t *testing.T) {

			if got := myAtoi(s.in); got != s.res {
				t.Errorf("not eq got %d, expected %d", got, s.res)
			}

		})

	}

}
