package zigzag

import "testing"

func TestConvert(t *testing.T) {

	tt := []struct {
		in   string
		rows int
		res  string
	}{{
		in:   "PINALSIGYAHRPI",
		rows: 4,
		res:  "P     I    N\nA   L S  I G\nY A   H R\nP     I",
	}}

	for _, s := range tt {
		t.Run(s.in, func(t *testing.T) {

			if got := convert(s.in, s.rows); got != s.res {
				t.Errorf("not eq got %s, expected %s", got, s.res)
			}

		})

	}

}
