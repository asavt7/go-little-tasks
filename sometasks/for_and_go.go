package main

import (
	"sync"
)

func main() {
	// 	no messages in stdout (all of )
	// runtime.GOMAXPROCS(1)

	wg := sync.WaitGroup{}
	max := 10000
	for i := 0; i < max; i++ {
		wg.Add(1)
		var j int = i
		go func() {
			var a int = i
			if a != max {
				println("i= ", i, " j=", j)
			}
			wg.Done()
		}()
	}
	wg.Wait()
}
