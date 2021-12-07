package longest

import "testing"

func TestLongest(t *testing.T) {
	tt := []struct {
		s string
		r int
	}{{
		s: "abcabcbb",
		r: 3,
	},
		{
			s: "bbbbb",
			r: 1,
		}, {
			s: "pwwkew",
			r: 3,
		}, {
			s: "",
			r: 0,
		},
		{
			s: " ",
			r: 1,
		},
	}

	for _, s := range tt {
		t.Run(s.s, func(t *testing.T) {
			if got := lengthOfLongestSubstring(s.s); got != s.r {
				t.Errorf("err expected %d got %d", s.r, got)
			}
		})
	}

}
