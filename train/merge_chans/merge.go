package merge

import (
	"reflect"
	"sync"
)

func MergeChannelsReflect(channels ...<-chan int) <-chan int {
	res := make(chan int)

	var cases []reflect.SelectCase
	for _, c := range channels {
		cases = append(cases, reflect.SelectCase{
			Dir:  reflect.SelectRecv,
			Chan: reflect.ValueOf(c),
		})
	}

	go func() {
		for len(cases) > 0 {
			i, v, ok := reflect.Select(cases)
			if !ok {
				cases = append(cases[:i], cases[i+1:]...)
				continue
			}
			res <- v.Interface().(int)
		}
		close(res)
	}()

	return res
}

func MergeChannelsGoroutines(channels ...<-chan int) <-chan int {
	res := make(chan int)

	wg := sync.WaitGroup{}
	wg.Add(len(channels))

	for _, ch := range channels {
		go func(ch <-chan int) {
			for v := range ch {
				res <- v
			}
			wg.Done()
		}(ch)
	}

	go func() {
		wg.Wait()
		close(res)
	}()

	return res
}

func createChan(numOfMessages int) <-chan int {
	c := make(chan int)

	go func() {
		for i := 0; i < numOfMessages; i++ {
			c <- i
		}
		close(c)
	}()
	return c
}
