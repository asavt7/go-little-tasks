package main

import (
	"sort"
)

func fourSum(nums []int, target int) [][]int {

	if len(nums) < 4 {
		return nil
	}

	nums = rmDubles(nums)

	res := make([][]int, 0)

	m := make(map[int][][]int)
	for i := 0; i < len(nums)-1; i++ {
		for j := i + 1; j < len(nums); j++ {
			if i == j {
				continue
			}
			s := nums[i] + nums[j]
			m[s] = append(m[s], []int{i, j})
		}
	}

	for k, _ := range m {
		if v, ok := m[target-k]; ok {
			for _, vals := range m[k] {
				for _, vals1 := range v {
					if vals1[0] == vals[0] || vals1[0] == vals[1] || vals1[1] == vals[0] || vals1[1] == vals[1] {
						continue
					}

					x := []int{nums[vals[0]], nums[vals[1]], nums[vals1[0]], nums[vals1[1]]}
					sort.Ints(x)
					if !containsArr(res, x) {
						res = append(res, x)
					}
				}

			}
		}

	}
	return res
}

func rmDubles(nums []int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	l := 0
	for k, v := range m {
		max := 4
		if v < 4 {
			max = v
		}
		for i := 0; i < max; i++ {
			nums[l] = k
			l++
		}
	}
	return nums[:l]
}
