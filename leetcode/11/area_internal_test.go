package area

import (
	"fmt"
	"testing"
)

func TestName(t *testing.T) {

	tt := []struct {
		in  []int
		res int
	}{{
		in:  []int{1, 1},
		res: 1,
	}, {
		in:  []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
		res: 49,
	},
	}

	for _, s := range tt {
		t.Run(fmt.Sprintf("%+v", s.in), func(t *testing.T) {

			if got := maxArea(s.in); got != s.res {
				t.Errorf("not eq got %d, expected %d", got, s.res)
			}

		})

	}

}
