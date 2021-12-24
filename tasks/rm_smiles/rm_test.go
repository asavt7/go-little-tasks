package rm_test

import (
	rm "github.com/asavt7/go-little-tasks/tasks/rm_smiles"
	"strings"
	"testing"
)

func TestRmSmiles(t *testing.T) {
	tt := []struct {
		in, exp string
	}{{in: "", exp: ""},
		{
			in:  ":-)",
			exp: "",
		},
		{
			in:  "::-)",
			exp: ":",
		},
		{
			in:  ":-)))",
			exp: "",
		},
		{
			in:  ":--)))",
			exp: ":--)))",
		},

		{
			in:  " :-)))1",
			exp: " 1",
		},
		{
			in:  ":-)))((",
			exp: "((",
		},
		{
			in:  "Я работаю в Яндексе :-)))",
			exp: "Я работаю в Яндексе ",
		},
		{
			in:  "везет :-) а я еще нет :-((",
			exp: "везет  а я еще нет ",
		},
	}

	for _, s := range tt {
		t.Run(s.in, func(t *testing.T) {
			if got := rm.RmSmiles(s.in); got != s.exp {
				t.Errorf("exp: %s != got: %s", s.exp, got)
			}
		})
	}
}

func BenchmarkRmSmiles(b *testing.B) {
	s := strings.Repeat(" везет :-) а я еще нет :-(( ", 100)

	for i := 0; i < b.N; i++ {
		_ = rm.RmSmiles(s)
	}
}

func BenchmarkRmSmiles_splits(b *testing.B) {
	s := strings.Repeat(" везет :-) а я еще нет :-(( ", 100)

	for i := 0; i < b.N; i++ {
		_ = rm.RmSmiles_splits(s)
	}
}
