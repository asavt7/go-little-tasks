package main

import "fmt"

type A struct {
	a string
	b int

	//	i []int //cannot use as key
	//	m map[struct{}]struct{} //cannot use as key
}

type B struct {
	a *A
	//	foo func() //cannot be key of map
}

func main() {

	m1 := make(map[A]B)

	for b, a := range m1 {
		fmt.Println(b, a)
	}

}
