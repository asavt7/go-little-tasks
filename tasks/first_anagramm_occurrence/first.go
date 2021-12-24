package first

func FirstAnagramOccurrence(instr string, substr string) int {
	if len(instr) == 0 {
		if len(substr) > 0 {
			return -1
		}
		return 0
	}

	in := []rune(instr)

	searching := make(map[rune]int)
	for _, i2 := range substr {
		searching[i2]++
	}
	l, r := 0, len(substr)-1

	m := make(map[rune]int)
	for _, i := range in[l : r+1] {
		m[i]++
	}

	found := true
	for k, _ := range searching {
		if m[k] != searching[k] {
			found = false
			break
		}
	}
	if found {
		return l
	}

	for r < len(in) {
		found = true
		for k, _ := range searching {
			if m[k] != searching[k] {
				found = false
				break
			}
		}
		if found {
			return l
		}

		m[in[l]]--
		l++
		r++
		if r == len(in) {
			return -1
		}
		m[in[r]]++
	}

	return -1
}
