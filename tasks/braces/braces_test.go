package braces_test

import (
	"github.com/asavt7/go-little-tasks/tasks/braces"
	"reflect"
	"strconv"
	"testing"
)

func TestGenerateCorrectSequences(t *testing.T) {
	tt := []struct {
		n        int
		expected []string
	}{{
		n:        1,
		expected: []string{"()"},
	},
		{
			n:        2,
			expected: []string{"(())", "()()"},
		},
		{
			n:        3,
			expected: []string{"((()))", "(()())", "(())()", "()(())", "()()()"},
		},
	}

	for i, s := range tt {
		t.Run(strconv.Itoa(i), func(t *testing.T) {
			got := braces.GenerateCorrectSequences(s.n)
			if !reflect.DeepEqual(got, s.expected) {
				t.Errorf("expected %+v\n got %+v", s.expected, got)
			}
		})
	}

}
