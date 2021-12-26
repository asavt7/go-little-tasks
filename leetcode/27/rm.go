package main

func removeElement(nums []int, val int) int {
	if len(nums) == 0 {
		return 0
	}
	if len(nums) == 1 {
		if nums[0] == val {
			return 0
		}
		return 1
	}
	r := len(nums)
	for r > 0 && nums[r-1] == val {
		r--
	}

	for i := 0; i < r; i++ {
		if nums[i] == val {
			r--
			nums[i] = nums[r]
			for r > 0 && nums[r-1] == val {
				r--
			}
			continue
		}
	}

	return r
}
