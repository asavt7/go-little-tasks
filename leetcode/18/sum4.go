package main

import "sort"

//timeout exceeded
func fourSum1(nums []int, target int) [][]int {

	if len(nums) < 4 {
		return nil
	}

	m := make(map[int][]int)

	res := make([][]int, 0)
	for i := 0; i < len(nums)-3; i++ {
		for j := i + 1; j < len(nums)-2; j++ {
			for e := j + 1; e < len(nums)-1; e++ {
				for w := e + 1; w < len(nums); w++ {
					if i == j || j == e || w == e {
						continue
					}
					if nums[i]+nums[j]+nums[e]+nums[w] == target {
						if v, ok := m[i]; ok {
							if contains(v, j) && contains(v, e) && contains(v, w) {
								continue
							}
						}
						m[i] = []int{j, e, w}
						x := []int{nums[i], nums[j], nums[e], nums[w]}
						sort.Ints(x)
						if !containsArr(res, x) {
							res = append(res, x)
						}
					}
				}
			}
		}

	}
	return res
}

func contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i, v := range a {
		if v != b[i] {
			return false
		}
	}
	return true
}

func containsArr(s [][]int, e []int) bool {
	for _, a := range s {
		if Equal(a, e) {
			return true
		}
	}
	return false
}
