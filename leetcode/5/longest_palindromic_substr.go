package leet5

//https://leetcode.com/problems/longest-palindromic-substring/

func longestPalindrome(s string) string {
	if len(s) == 1 {
		return s
	}

	res := ""
	for i := 0; i < len(s)-1; i++ {
		cur := check(s, i)
		if len(cur) > len(res) {
			res = cur
		}
		cur = checkTwo(s, i, i+1)
		if len(cur) > len(res) {
			res = cur
		}
	}
	return res
}

func check(s string, i int) string {
	left, right := i, i
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return s[left+1 : right]
}

func checkTwo(s string, i, j int) string {
	left, right := i, j
	for left >= 0 && right < len(s) && s[left] == s[right] {
		left -= 1
		right += 1
	}
	return s[left+1 : right]
}
