package median

func findMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	res := make([]int, len(nums1)+len(nums2))

	cur := 0
	i1 := 0
	i2 := 0

	for i1 < len(nums1) && i2 < len(nums2) {
		if nums1[i1] < nums2[i2] {
			res[cur] = nums1[i1]
			i1++
		} else {
			res[cur] = nums2[i2]
			i2++
		}
		cur++
	}
	for i1 < len(nums1) {
		res[cur] = nums1[i1]
		i1++
		cur++
	}
	for i2 < len(nums2) {
		res[cur] = nums2[i2]
		i2++
		cur++
	}

	if len(res) == 0 {
		return 0
	}

	if len(res)%2 == 0 {
		return float64(res[len(res)/2]+res[len(res)/2-1]) / 2.
	} else {
		return float64(res[len(res)/2])
	}
}
