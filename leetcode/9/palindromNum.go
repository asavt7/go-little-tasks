package pal

import (
	"strconv"
)

func isPalindrome(x int) bool {
	return isPalindromeStr(x)
}

func isPalindromeStr(x int) bool {
	s := strconv.Itoa(x)
	var l, r = 0, len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}

func isPalindromeNum(x int) bool {
	s := make([]int, 0, 19)
	for x > 0 {
		s = append(s, x%10)
		x = x / 10
	}

	var l, r = 0, len(s) - 1
	for l < r {
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}
