package main

import "fmt"

func strStr(h string, n string) int {
	if len(h) < len(n) {
		return -1
	}

	l, r := 0, len(n)-1

	f := true
	for r < len(h) {
		f = true
		for i := 0; i < len(n); i++ {
			if h[l+i] != n[i] || h[r-i] != n[len(n)-1-i] {
				f = false
				if i == 0 {
					i = 1
				}
				r = r + i
				l = l + i
				break
			}
		}
		if f {
			return l
		}

	}

	return -1
}

func main() {
	fmt.Println(strStr("mississippi", "issip"))
}
