package merge

import "sort"

func merge(intervals [][]int) [][]int {
	return Merge(intervals)
}

func Merge(intervals [][]int) [][]int {
	if len(intervals) < 2 {
		return intervals
	}

	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	var last, cur int = 0, 1

	for ; cur < len(intervals); cur++ {
		if intervals[last][1] >= intervals[cur][0] {
			if intervals[last][1] < intervals[cur][1] {
				intervals[last][1] = intervals[cur][1]
			}
			continue
		}

		last++
		intervals[last] = intervals[cur]
	}

	return intervals[:last+1]
}
