package main

import "fmt"

func calc() func(int) int {
	f := 0
	return func(i int) int {
		f += i / 2
		return f
	}
}

func main() {
	a, b := calc(), calc()

	for i := 0; i < 3; i++ {
		fmt.Println(a(i), b(-3*i))
	}
}
