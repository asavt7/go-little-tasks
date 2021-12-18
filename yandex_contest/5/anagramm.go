package main

import "fmt"

func main() {
	var a, b string

	fmt.Scan(&a)
	fmt.Scan(&b)

	m := make(map[rune]int)
	for _, v := range a {
		m[v]++
	}
	for _, v := range b {
		_, ok := m[v]
		if !ok {
			fmt.Println("0")
			return
		}
		m[v]--
	}

	for _, v := range m {
		if v != 0 {
			fmt.Println("0")
			return
		}
	}

	fmt.Println("1")
	return
}
