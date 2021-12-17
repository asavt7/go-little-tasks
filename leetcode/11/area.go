package area

func maxArea(height []int) int {
	best := 0
	curr := 0

	for l, r := 0, len(height)-1; l < r; {
		h := min(height[l], height[r])
		w := r - l
		curr = w * h
		if curr > best {
			best = curr
		}
		if height[l] < height[r] {
			l++
		} else {
			r--
		}
	}

	return best
}

func min(a, b int) int {
	if a < b {
		return a
	} else {
		return b
	}
}
