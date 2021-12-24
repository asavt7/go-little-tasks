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

func FirstAnagramOccurrence_OptHash(instr string, substr string) int {
	if len(instr) == 0 {
		if len(substr) > 0 {
			return -1
		}
		return 0
	}

	in := []rune(instr)
	searchingHash := hash([]rune(substr))

	searching := make(map[rune]int)
	for _, i2 := range substr {
		searching[i2]++
	}
	l, r := 0, len(substr)-1

	m := make(map[rune]int)
	for _, i := range in[l : r+1] {
		m[i]++
	}

	substrHash := hash(in[l : r+1])

	found := true
	if substrHash == searchingHash {
		for k, _ := range searching {
			if m[k] != searching[k] {
				found = false
				break
			}
		}
		if found {
			return l
		}
	}

	for r < len(in) {
		if substrHash == searchingHash {
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
		}

		substrHash -= in[l]
		m[in[l]]--
		l++
		r++
		if r == len(in) {
			return -1
		}
		m[in[r]]++
		substrHash += in[r]
	}

	return -1
}

func hash(s []rune) int32 {
	var r int32
	for _, i2 := range s {
		r += i2
	}
	return r
}
