package first_test

import (
	first "github.com/asavt7/go-little-tasks/tasks/first_anagramm_occurrence"
	"testing"
)

func TestFirstAnagramOccurrence(t *testing.T) {

	tt := []struct {
		in, substr string
		exp        int
	}{
		{in: "", substr: ""},
		{
			in:     "abcdeqwertyui",
			substr: "cdb",
			exp:    1,
		},
		{
			in:     "abcdeqwertyui",
			substr: "iuy",
			exp:    10,
		},
		{
			in:     "abcdeqwertyui",
			substr: "i",
			exp:    12,
		},
		{
			in:     "abcdeqwertyuiabcd",
			substr: "bcd",
			exp:    1,
		},
	}

	for _, s := range tt {
		t.Run(s.in, func(t *testing.T) {
			if got := first.FirstAnagramOccurrence(s.in, s.substr); got != s.exp {
				t.Errorf("%+v substr: %d != got: %d", s, s.exp, got)
			}
		})
	}

}
