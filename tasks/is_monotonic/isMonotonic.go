package main

// {1,7,9} - true
// {1,1,1} - true
// {9,5,1} - true
// {1,2,3,4,5,4} - false
// {23,5,23} - false
// {1,5,1} - false
// {} - true
// {1,2} - true
// {1,1} - true
// {2,1} - true

func isMonotonic(in []int) bool {
	if len(in) < 3 {
		return true
	}

	// признак массив возрастает
	isUp := false
	if in[0] <= in[len(in)-1] {
		isUp = true
	}

	prev := in[0]
	for i := 1; i < len(in); i++ {
		if isUp {
			if in[i] < prev {
				return false
			}
		} else {
			if in[i] > prev {
				return false
			}
		}
		prev = in[i]
	}
	return true
}

func isMonotonic1(in []int) bool {
	if len(in) < 3 {
		return true
	}
	isUp, isDown := true, true

	for i := 1; i < len(in); i++ {
		isUp = isUp && in[i-1] <= in[i]
		isDown = isDown && in[i-1] >= in[i]
	}
	return isDown || isUp
}

func main() {
	println(isMonotonic1([]int{1, 2, 1}))
	println(isMonotonic1([]int{1, 2, 3}))
	println(isMonotonic1([]int{3, 2, 1}))
	println(isMonotonic1([]int{3, 2, 1, 2, 3, 4}))
	println(isMonotonic1([]int{1, 1, 1, 1}))

}
