package main

import "fmt"

func main() {

	var n int
	fmt.Scanf("%d", &n)
	generate(n)

}

func generate(n int) {
	gen("", 0, 0, n)

}

func gen(cur string, left, right, n int) {
	if len(cur) == 2*n {
		fmt.Println(cur)
		return
	}

	if left < n {
		gen(cur+"(", left+1, right, n)
	}
	if right < left {
		gen(cur+")", left, right+1, n)
	}
}
