package intervals

import "sort"

type A struct {
	In, Out int
}

func MaxCap(s []A) int {
	if len(s) < 2 {
		return len(s)
	}

	var startAt []int
	var endAt []int

	for _, v := range s {
		startAt = append(startAt, v.In)
		endAt = append(endAt, v.Out)
	}

	sort.Ints(startAt)
	sort.Ints(endAt)

	l, r := 0, 0
	maxC, c := 0, 0

	for l < len(s) && r < len(s) {
		if startAt[l] < endAt[r] {
			c++
			maxC = max(c, maxC)
			l++
		} else {
			c--
			r++
		}
	}

	return maxC
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
