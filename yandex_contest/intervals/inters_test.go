package intervals_test

import (
	"github.com/asavt7/go-little-tasks/yandex_contest/intervals"
	"strconv"
	"testing"
)

func TestMaxCap(t *testing.T) {

	tt := []struct {
		in       []intervals.A
		expected int
	}{
		{
			in:       []intervals.A{},
			expected: 0,
		},
		{
			in: []intervals.A{
				{In: 0, Out: 5},
			},
			expected: 1,
		},
		{
			in: []intervals.A{
				{In: 0, Out: 5},
				{In: 0, Out: 5},
				{In: 1, Out: 5},
				{In: 4, Out: 6},
				{In: 5, Out: 9},
				{In: 8, Out: 9},
			},
			expected: 4,
		},
		{
			in: []intervals.A{
				{In: 0, Out: 5},
				{In: 0, Out: 5},
				{In: 1, Out: 5},
				{In: 4, Out: 6},
				{In: 5, Out: 9},
				{In: 0, Out: 1},
			},
			expected: 4,
		},
		{
			in: []intervals.A{
				{In: 0, Out: 5},
				{In: 0, Out: 5},
				{In: 1, Out: 5},
				{In: 4, Out: 6},
				{In: 5, Out: 9},
				{In: 0, Out: 1},
				{In: 0, Out: 1},
				{In: 0, Out: 1},
				{In: 0, Out: 1},
			},
			expected: 6,
		},
		{
			in: []intervals.A{
				{In: 0, Out: 5},
				{In: 10, Out: 15},
				{In: 0, Out: 15},
			},
			expected: 2,
		},
		{
			in: []intervals.A{
				{In: 0, Out: 5},
				{In: 10, Out: 15},
				{In: 100, Out: 150},
			},
			expected: 1,
		},
		{
			in: []intervals.A{
				{In: 0, Out: 5},
				{In: 0, Out: 15},
				{In: 0, Out: 150},
			},
			expected: 3,
		},
		{
			in: []intervals.A{
				{In: 0, Out: 150},
				{In: 0, Out: 1},
				{In: 0, Out: 1},
				{In: 0, Out: 1},
				{In: 0, Out: 1},
				{In: 149, Out: 150},
				{In: 149, Out: 150},
				{In: 149, Out: 150},
				{In: 149, Out: 150},
			},
			expected: 5,
		},
	}

	for i, s := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			if actual := intervals.MaxCap(s.in); s.expected != actual {
				t.Errorf("expected %d actual %d\n", s.expected, actual)
			}
		})
	}
}
