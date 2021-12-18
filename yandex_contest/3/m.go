package main

import (
	"fmt"
)

func main() {
	var count, c, prev int32
	var res []int32
	fmt.Scanf("%d", &count)

	if count < 1 {
		return
	}
	fmt.Scanf("%d", &prev)
	fmt.Printf("%d\n", prev)

	var i int32 = 1
	for ; i < count; i++ {
		fmt.Scanf("%d", &c)
		if c > prev {
			res = append(res, c)
			prev = c
		}
	}

	for _, c := range res {
		fmt.Printf("%d\n", c)
	}

}
