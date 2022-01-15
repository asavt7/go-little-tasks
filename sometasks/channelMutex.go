package main

import (
	"fmt"
	"sync"
)

type MyMutex struct {
	lockCh chan struct{}
}

func NewMyMutex() *MyMutex {
	ch := make(chan struct{})
	go func() {
		ch <- struct{}{}
	}()
	return &MyMutex{lockCh: ch}
}

func (m *MyMutex) Unlock() {
	m.lockCh <- struct{}{}
}

func (m *MyMutex) Lock() {
	<-m.lockCh
}

func main() {
	c := 0
	m := NewMyMutex()

	wg := sync.WaitGroup{}

	for i := 0; i < 5000; i++ {
		wg.Add(1)
		go func() {
			m.Lock()
			defer m.Unlock()
			defer wg.Done()
			c++
		}()
	}
	wg.Wait()
	fmt.Println(c)
}
