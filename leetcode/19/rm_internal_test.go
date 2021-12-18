package main

import (
	"fmt"
	"reflect"
	"testing"
)

// Test_removeNthFromEnd comment
func Test_removeNthFromEnd(t *testing.T) {

	tt := []struct {
		in  []int
		n   int
		exp []int
	}{
		{
			in:  []int{1, 2, 3},
			n:   1,
			exp: []int{1, 2},
		},
		{
			in:  []int{1},
			n:   1,
			exp: []int{},
		},
		{
			in:  []int{1, 2},
			n:   2,
			exp: []int{2},
		},
		{
			in:  []int{1, 2},
			n:   1,
			exp: []int{1},
		},
		{
			in:  []int{1, 2, 3},
			n:   2,
			exp: []int{1, 3},
		},
	}

	for _, s := range tt {
		t.Run(fmt.Sprintf("%+v %d", s.in, s.n), func(t *testing.T) {
			in := generateNodes(s.in)
			got := removeNthFromEnd(in, s.n)
			exp := s.exp
			if !reflect.DeepEqual(got.toSlice(), exp) {
				t.Errorf("failed! got %+v exp %+v", got, exp)
			}

		})
	}

}
