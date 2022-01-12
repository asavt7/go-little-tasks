package main

import (
	"fmt"
	"sort"
)

func main() {
	d := [6]int{1, 2, 3, 4, 5}
	sd := d[2:6]
	sort.Ints(sd)
	fmt.Println(sd)
	// [0 3 4 5]
}
