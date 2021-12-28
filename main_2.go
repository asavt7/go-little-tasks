package main

import (
	"fmt"
)

func main() {

	fmt.Println(len([]rune("aфыфвa")))

	for _, r := range []byte("фыфвa") {
		fmt.Println(r)
	}

}

func foooooo(a []int) []int {

	for i := 0; i < 10; i++ {
		a = append(a, i)
	}
	fmt.Printf("inside foo : %+v\n", a)

	return a
}

func cal(a, b int) (c, d int) {

	defer func(a, b int) {
		c, d = a, b
	}(a, b)
	a, b = 33, 33
	c, d = a, b

	return 1111, 2222
}
