package main

import (
	"reflect"
	"testing"
)

// Test_removeNthFromEnd comment
func Test_removeNthFromEnd(t *testing.T) {

	in := generateNodes([]int{1, 2, 3})

	t.Run(t.Name(), func(t *testing.T) {
		got := removeNthFromEnd(in, 1)
		exp := []int{1, 2}
		if !reflect.DeepEqual(got.toSlice(), exp) {
			t.Errorf("failed! got %+v exp %+v", got, exp)
		}
	})

	t.Run(t.Name(), func(t *testing.T) {
		got := removeNthFromEnd(in, 1)
		exp := []int{1, 2, 3, 4}
		if !reflect.DeepEqual(got.toSlice(), exp) {
			t.Errorf("failed! got %+v exp %+v", got, exp)
		}

	})

}
