package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("--------- CATCH PANIC IN MAIN ------", err)
		}
	}()

	wg := sync.WaitGroup{}
	for i := 0; i < 10; i++ {

		wg.Add(1)
		go worker(&wg)()
	}

	wg.Wait()

	fmt.Println("--------------- END --------------")
}

func worker(wg *sync.WaitGroup) func() {

	return func() {

		defer wg.Done()
		time.Sleep(100 * time.Millisecond)

		//		if rand.Intn(2) == 1 {
		panic("from worker panic")
		//		}
		fmt.Println("do work")

	}
}
