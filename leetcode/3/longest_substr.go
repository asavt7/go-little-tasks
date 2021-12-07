package longest

//https://leetcode.com/problems/longest-substring-without-repeating-characters/

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]struct{})
	l := 0
	cur := 0
	for i := 0; i < len(s); i++ {

		cur = 0
		j := i
		for ; j < len(s); j++ {
			if _, ok := m[s[j]]; ok {
				if cur > l {
					l = cur
				}
				break
			}
			m[s[j]] = struct{}{}
			cur++
		}
		if j == len(s) {
			if cur > l {
				l = cur
			}
			break
		}
		m = make(map[byte]struct{})
	}

	return l
}
