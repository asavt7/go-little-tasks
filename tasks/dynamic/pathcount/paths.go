package main

import (
	"fmt"
)

// https://www.youtube.com/watch?v=GhiRlhPlJ9Q

func Paths(m, n int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	//start at point (1,1)
	if m == 1 || n == 1 {
		return 1
	}

	return Paths(m-1, n) + Paths(m, n-1)
}

func PathsDynamic(m, n int) int {
	arr := make([][]int, m)
	for i := 0; i < len(arr); i++ {
		arr[i] = make([]int, n)
	}
	return pathsHelper(m, n, arr)
}

func pathsHelper(m, n int, arr [][]int) int {
	if m <= 0 || n <= 0 {
		return 0
	}
	//start at point (1,1)
	if m == 1 || n == 1 {
		return 1
	}

	if arr[m-1][n-1] != 0 {
		return arr[m-1][n-1]
	}

	arr[m-1][n-1] = pathsHelper(m-1, n, arr) + pathsHelper(m, n-1, arr)
	return arr[m-1][n-1]
}

func main() {
	fmt.Println(Paths(4, 5))
	fmt.Println(PathsDynamic(4, 5))

}
