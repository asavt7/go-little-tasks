package main

import (
	"fmt"
	"sync"
)

// fatal error: concurrent map writes !!!
func main() {
	m := make(map[int]int64, 100256)

	wg := sync.WaitGroup{}
	for i := 0; i < 256; i++ {
		wg.Add(1)
		go func(num int) {
			for i := 0; i < 100500; i++ {
				m[num]++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < 256; i++ {
		fmt.Printf("m[%d]=%d", i, m[i])
	}

}
