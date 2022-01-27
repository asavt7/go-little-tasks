package main

import "math"

func NumOfBracesRecursion(n int) int {
	if n < 1 {
		return 0
	}
	res := 0
	s := ""
	l, r := 0, 0

	perm(s, l, r, n, &res)

	return res
}

func perm(s string, l int, r int, n int, res *int) {
	if len(s) >= 2*n {
		*res++
		println(s)
		return
	}
	if l < n {
		perm(s+"(", l+1, r, n, res)
	}

	if r < l {
		perm(s+")", l, r+1, n, res)
	}
}

type BraceState struct {
	l, r int
}

func NumOfBraces(n int) int {
	if n < 1 {
		return 0
	}
	if n < 2 {
		return 1
	}
	result := 0

	//stack
	s := []BraceState{}
	s = append(s, BraceState{l: 1})
	c := 0
	for len(s) > 0 {
		c++
		pop := s[len(s)-1]
		s = s[:len(s)-1]
		if pop.r+pop.l == 2*n {
			result++
			continue
		}
		if pop.l < n {
			s = append(s, BraceState{
				l: pop.l + 1,
				r: pop.r,
			})
		}
		if pop.r < pop.l {
			s = append(s, BraceState{
				l: pop.l,
				r: pop.r + 1,
			})
		}
	}

	println("count of stack pop operations", c)
	return result
}

func main() {
	for i := 2; i < 100; i++ {
		a := NumOfBraces(i)
		b := math.Pow(float64(i), math.Log(float64(i)))
		println(i, a, int(b), int(b)-a)
	}

}
