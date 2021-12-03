package anagramms

func IsAnagrams(s1, s2 string) bool {
	if len(s1) != len(s2) {
		return false
	}

	m := make(map[rune]int)
	for _, i2 := range s1 {
		m[i2] += 1
	}
	for _, i2 := range s2 {
		m[i2] -= 1
	}
	for _, val := range m {
		if val != 0 {
			return false
		}
	}

	return true
}
