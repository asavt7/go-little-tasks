package sortlinkedlist_test

import (
	"fmt"
	sortlinkedlist "github.com/asavt7/go-little-tasks/leetcode/148"
	. "github.com/asavt7/go-little-tasks/leetcode/utils"
	"reflect"
	"testing"
)

func Test_(t *testing.T) {

	tt := []struct {
		in  []int
		exp []int
	}{
		{
			in:  []int{5, 4, 3, 2, 1},
			exp: []int{1, 2, 3, 4, 5},
		},
		{
			in:  []int{1, 2, 3},
			exp: []int{1, 2, 3},
		},
		{
			in:  []int{1},
			exp: []int{1},
		},
		{
			in:  []int{2, 1, 1},
			exp: []int{1, 1, 2},
		},
		{
			in:  []int{4, 2, 1, 3},
			exp: []int{1, 2, 3, 4},
		},
	}

	for _, s := range tt {
		t.Run(fmt.Sprintf("in %+v \nexp %+v", s.in, s.exp), func(t *testing.T) {
			in := GenerateNodes(s.in)
			got := sortlinkedlist.SortList(in)
			exp := s.exp
			if !reflect.DeepEqual(got.ToSlice(), exp) {
				t.Errorf("failed! got %+v exp %+v", got, exp)
			}
		})
	}

}
