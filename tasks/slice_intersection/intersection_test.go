package intersection_test

import (
	intersection "github.com/asavt7/go-little-tasks/tasks/slice_intersection"
	"reflect"
	"sort"
	"strconv"
	"testing"
)

func TestIntersection(t *testing.T) {
	tt := []struct {
		c1       []int
		c2       []int
		expected []int
	}{
		{
			c1:       []int{1, 2},
			c2:       []int{1, 2},
			expected: []int{1, 2},
		},
		{
			c1:       []int{},
			c2:       []int{1, 2},
			expected: []int{},
		},
		{
			c1:       []int{1, 2},
			c2:       []int{1},
			expected: []int{1},
		},
		{
			c1:       []int{1, 2, 3},
			c2:       []int{4, 5, 6},
			expected: []int{},
		},
	}

	for i, s := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			inter := intersection.Intersection(s.c1, s.c2)

			sort.Ints(inter)

			if !reflect.DeepEqual(s.expected, inter) {
				t.Errorf("not equal\nexpected : %v\nactual : %v", s.expected, inter)
			}
		})
	}

}
