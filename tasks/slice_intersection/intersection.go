package intersection

func Intersection(c1, c2 []int) []int {
	inters := make(map[int]int)
	res := make([]int, 0)

	for _, e := range c1 {
		inters[e] += 1
	}

	for _, e := range c2 {
		if _, ok := inters[e]; ok {
			inters[e] -= 1
			res = append(res, e)
		}
	}
	return res
}
