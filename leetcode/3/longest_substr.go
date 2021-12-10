package longest

//https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	if len(s) < 2 {
		return len(s)
	}

	m := make(map[byte]bool)
	longest := 0
	cur := 0
	left, right := 0, 0

	for right < len(s) {
		if v, _ := m[s[right]]; v {
			if cur > longest {
				longest = cur
			}
			cur--
			m[s[left]] = false
			left++
			continue
		}

		m[s[right]] = true
		cur++
		right++
	}
	if cur > longest {
		longest = cur
	}

	return longest
}
