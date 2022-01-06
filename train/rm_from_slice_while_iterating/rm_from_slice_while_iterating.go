package main

import "fmt"

func main() {
	a := make([]int, 100)
	for i := 0; i < len(a); i++ {
		a[i] = i
	}

	// remove every tenth element
	for i, removed, j := 0, 0, 0; i < len(a); i++ {
		j = i - removed
		if i%10 == 0 {
			a = append(a[:j], a[j+1:]...)
			removed++
		}
	}

	fmt.Println(len(a))
}
