package main

import "fmt"

//https://www.youtube.com/watch?v=-59FbGWsCgI

func main() {
	fmt.Println(temps([]int{13, 12, 15, 11, 9, 12, 16}))
	fmt.Println(tempsStack([]int{13, 12, 15, 11, 9, 12, 16}))

}

func temps(in []int) []int {
	s := make([]int, 0, len(in))
	result := make([]int, len(in))

	for i := 0; i < len(in)-1; i++ {
		s = s[:0]
		max := in[i]
		for j := i + 1; j < len(in); j++ {
			s = append(s, in[j])
			if in[j] > max {
				result[i] = len(s)
				break
			}
		}
	}

	return result
}

func tempsStack(in []int) []int {
	s := make([]int, 0, len(in))
	result := make([]int, len(in))

	for i := len(in) - 1; i >= 0; i-- {
		for len(s) > 0 && in[i] > in[s[len(s)-1]] {
			s = s[:len(s)-1]
		}

		if len(s) == 0 {
			s = append(s, i)
			continue
		}
		result[i] = s[len(s)-1] - i
		s = append(s, i)
	}
	return result
}
