package median

import (
	"math"
	"testing"
)

func TestMedian(t *testing.T) {

	tt := []struct {
		in1 []int
		in2 []int
		exp float64
	}{{
		in1: []int{1, 2},
		in2: []int{3, 4},
		exp: 2.5,
	}, {
		in1: []int{},
		in2: []int{2, 3},
		exp: 2.5,
	}}

	for _, s := range tt {
		t.Run("", func(t *testing.T) {
			if got := findMedianSortedArrays(s.in1, s.in2); math.Abs(s.exp-got) > 0.0001 {
				t.Errorf("exp %v but got %v", s.exp, got)
			}
		})
	}
}
