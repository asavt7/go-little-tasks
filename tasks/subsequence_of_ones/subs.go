package subsequence_of_ones

func MaxSubsequenceLength(nums []int) int {
	best := 0
	current := 0

	for _, num := range nums {
		if num > 0 {
			current += 1
		} else {
			if current > best {
				best = current
			}
			current = 0
		}
	}
	if current > best {
		return current
	}
	return best

}
