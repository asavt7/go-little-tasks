package main

import "fmt"

const (
	a = 1
	b = 2
	c
	d
	e = 5
)

func main() {
	fmt.Println(a, b, c, d, e)
	//c,d == 2
}
