package leet5

import "testing"

//https://leetcode.com/problems/longest-palindromic-substring/

func TestCheck(t *testing.T) {
	tt := []struct{ in, exp string }{{
		in:  "babad",
		exp: "bab",
	}, {
		in:  "cbbd",
		exp: "bb",
	}}

	for _, s := range tt {
		t.Run(s.in, func(t *testing.T) {
			if got := longestPalindrome(s.in); s.exp != got {
				t.Errorf("exp %s but got %s", s.exp, got)
			}
		})
	}
}
