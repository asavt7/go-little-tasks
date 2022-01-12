package main

import "fmt"

func swap(int, int) (x int, y int) {
	fmt.Println("swap", x, y)
	x, y = y, x
	fmt.Println("after swap", x, y)

	return
}

func main() {
	a, b := 1, 2
	a, b = swap(a, b)
	fmt.Println(a, b)

	var d []int
	d = append(d, 0)
	d = append(d, 1)
	d = append(d, 2)
	d[0], d[1] = swap(d[0], d[1])
	x := len(d) - 1
	fmt.Println(d[:x])
}
