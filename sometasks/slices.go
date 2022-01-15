package main

import "fmt"

func main() {

	a := []int{1, 2, 3}
	fmt.Println("before", a, len(a), cap(a))
	appendSliceFoo(a)
	fmt.Println("after appends", a, len(a), cap(a))
	changeSliceFoo(a)
	fmt.Println("after change", a, len(a), cap(a))

}

func changeSliceFoo(ints []int) {
	if len(ints) < 1 {
		return
	}
	ints[0] = 1234
}

func appendSliceFoo(ints []int) {
	ints = append(ints, 9999)
	ints = append(ints, 9999)

}
