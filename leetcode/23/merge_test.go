package merge_test

import (
	"fmt"
	merge "github.com/asavt7/go-little-tasks/leetcode/23"
	"reflect"
	"testing"
)

func TestMerge(t *testing.T) {

	tt := []struct {
		in  [][]int
		exp []int
	}{
		{
			in:  [][]int{{1, 2, 3}, {1, 2, 3}},
			exp: []int{1, 1, 2, 2, 3, 3},
		},
		{
			in:  [][]int{{}, {}},
			exp: []int{},
		},
		{
			in:  [][]int{},
			exp: []int{},
		},
		{
			in:  [][]int{{1}, {}},
			exp: []int{1},
		},
	}

	for _, s := range tt {
		t.Run(fmt.Sprintf("%+v", s.in), func(t *testing.T) {
			in := merge.GenerateNodeLists(s.in)
			got := merge.Merge(in)
			exp := s.exp
			if !reflect.DeepEqual(got.ToSlice(), exp) {
				t.Errorf("failed! got %+v exp %+v", got, exp)
			}

		})
	}

}
