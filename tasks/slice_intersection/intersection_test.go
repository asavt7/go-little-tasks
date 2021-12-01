package intersection_test

import (
	intersection "github.com/asavt7/go-little-tasks/tasks/slice_intersection"
	"reflect"
	"sort"
	"strconv"
	"strings"
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

type A struct {
	name string
}

func TestIntersectionInterface(t *testing.T) {
	tt := []struct {
		c1       []interface{}
		c2       []interface{}
		expected []interface{}
	}{
		{
			c1:       []interface{}{A{name: "a"}},
			c2:       []interface{}{A{name: "a"}},
			expected: []interface{}{A{name: "a"}},
		},
		{
			c1:       []interface{}{A{name: "a"}},
			c2:       []interface{}{A{name: "b"}},
			expected: []interface{}{},
		},
		{
			c1:       []interface{}{A{name: "a"}},
			c2:       []interface{}{A{name: "b"}, A{name: "a"}},
			expected: []interface{}{A{name: "a"}},
		},
		{
			c1:       []interface{}{A{name: "b"}, A{name: "a"}},
			c2:       []interface{}{A{name: "b"}, A{name: "a"}},
			expected: []interface{}{A{name: "a"}, A{name: "b"}},
		},
		{
			c1:       []interface{}{A{name: "a"}, A{name: "b"}, A{name: "b"}},
			c2:       []interface{}{A{name: "b"}, A{name: "a"}, A{name: "b"}},
			expected: []interface{}{A{name: "a"}, A{name: "b"}, A{name: "b"}},
		},
	}

	for i, s := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			inter := intersection.IntersectionInterface(s.c1, s.c2)

			sort.Slice(inter, func(i, j int) bool {
				return strings.Compare(inter[j].(A).name, inter[i].(A).name) > 0
			})

			if !reflect.DeepEqual(s.expected, inter) {
				t.Errorf("not equal\nexpected : %v\nactual : %v", s.expected, inter)
			}
		})
	}
}

func TestIntersectionInterface_strings(t *testing.T) {
	tt := []struct {
		c1       []interface{}
		c2       []interface{}
		expected []interface{}
	}{
		{
			c1:       []interface{}{"1"},
			c2:       []interface{}{"1"},
			expected: []interface{}{"1"},
		},
		{
			c1:       []interface{}{"1", "1", "1", "1"},
			c2:       []interface{}{"1", "1", "1"},
			expected: []interface{}{"1", "1", "1"},
		},
		{
			c1:       []interface{}{"1", "1", "1", "1", "2", "3"},
			c2:       []interface{}{"1", "1", "1", "2", "3"},
			expected: []interface{}{"1", "1", "1", "2", "3"},
		},
	}

	for i, s := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			inter := intersection.IntersectionInterface(s.c1, s.c2)

			sort.Slice(inter, func(i, j int) bool {
				return strings.Compare(inter[j].(string), inter[i].(string)) > 0
			})

			if !reflect.DeepEqual(s.expected, inter) {
				t.Errorf("not equal\nexpected : %v\nactual : %v", s.expected, inter)
			}
		})
	}

}
