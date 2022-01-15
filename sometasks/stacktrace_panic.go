package main

import (
	"fmt"
	"runtime/debug"
	"time"
)

func aa() {
	bb()
}

func bb() {
	fmt.Println("")
	cc()
}

func cc() {
	defer func() { fmt.Println("defer 3") }()
	defer func() { fmt.Println("defer 2") }()
	defer func() { fmt.Println("defer 1") }()

	panic("from cc func")
	debug.PrintStack()
}

func main() {

	go aa()
	time.Sleep(1 * time.Second)
}
