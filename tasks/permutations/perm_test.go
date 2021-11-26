package permutations_test

import (
	"fmt"
	"github.com/asavt7/go-little-tasks/tasks/permutations"
	"testing"
)

func TestPerm(t *testing.T) {
	table := []struct {
		inputData string
		expected  []string
	}{
		{
			inputData: "1",
			expected:  []string{"1"},
		},
		{
			inputData: "12",
			expected:  []string{"12", "21"},
		},
	}

	for i, s := range table {
		t.Run(fmt.Sprintf("test %d", i), func(t *testing.T) {
			resultSet := make(map[string]struct{})

			permutations.Perm([]rune(s.inputData), func(runes []rune) {
				resultSet[string(runes)] = struct{}{}
			})

			for _, q := range s.expected {
				if _, ok := resultSet[q]; !ok {
					t.Errorf("result incorrect : expected %+v , got : %+v", s.expected, resultSet)
				}
			}

		})
	}
}
