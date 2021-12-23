package main

import "fmt"

//26. Remove Duplicates from Sorted Array
func removeDuplicates(nums []int) int {
	if len(nums) < 2 {
		return len(nums)
	}
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[count-1] {
			nums[count] = nums[i]
			count++
		}
	}
	nums = nums[:count]
	return count
}

func main() {
	a := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(a))
	fmt.Println(a)
}
