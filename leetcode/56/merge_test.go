package merge_test

import (
	"fmt"
	merge "github.com/asavt7/go-little-tasks/leetcode/56"
	"reflect"
	"testing"
)

// Test_removeNthFromEnd comment
func Test_removeNthFromEnd(t *testing.T) {

	tt := []struct {
		in  [][]int
		exp [][]int
	}{
		{
			in:  [][]int{{1, 4}, {4, 5}},
			exp: [][]int{{1, 5}},
		},
		{
			in:  [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			exp: [][]int{{1, 6}, {8, 10}, {15, 18}},
		}}

	for _, s := range tt {
		t.Run(fmt.Sprintf("%+v %d", s.in, s.exp), func(t *testing.T) {
			if got := merge.Merge(s.in); !reflect.DeepEqual(got, s.exp) {
				t.Errorf("failed! got %+v exp %+v", got, s.exp)
			}
		})
	}
}
