package main

import (
	"fmt"
	"sync"
)

// ok , no fatal errors
func main() {
	arr := make([]int, 256, 256)

	wg := sync.WaitGroup{}
	for i := 0; i < 256; i++ {
		wg.Add(1)
		go func(num int) {
			for i := 0; i < 100500; i++ {
				arr[num]++
			}
			wg.Done()
		}(i)
	}
	wg.Wait()

	for i := 0; i < 256; i++ {
		fmt.Printf("m[%d]=%d\n", i, arr[i])
	}
}
