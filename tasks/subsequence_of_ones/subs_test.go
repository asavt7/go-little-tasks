package subsequence_of_ones_test

import (
	"github.com/asavt7/go-little-tasks/tasks/subsequence_of_ones"
	"strconv"
	"testing"
)

func TestMaxSubsequenceLength(t *testing.T) {
	tt := []struct {
		in       []int
		expected int
	}{
		{
			in:       []int{1, 1, 1, 1, 1, 0, 0, 0, 0},
			expected: 5,
		},
		{
			in:       []int{1, 1, 0, 1, 1, 0, 0, 0, 1},
			expected: 2,
		},
		{
			in:       []int{0, 0, 0},
			expected: 0,
		},
		{
			in:       []int{1},
			expected: 1,
		},
	}

	for i, s := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			actual := subsequence_of_ones.MaxSubsequenceLength(s.in)
			if s.expected != actual {
				t.Errorf("expected : %d but got : %d", s.expected, actual)
			}
		})
	}
}
