package stack

import (
	"fmt"
)

var EmptyStackError = fmt.Errorf("empty stack")

type Stack interface {
	Push(interface{})
	Pop() (interface{}, error)
	Top() (interface{}, error)
	Size() interface{}
}

type MyStack struct {
	arr []interface{}
}

func NewMyStack() *MyStack {
	return &MyStack{arr: make([]interface{}, 0)}
}

func (m *MyStack) Push(i interface{}) {
	m.arr = append(m.arr, i)
}

func (m *MyStack) Pop() (res interface{}, err error) {
	if len(m.arr) > 0 {
		res, m.arr = m.arr[len(m.arr)-1], m.arr[:len(m.arr)-1]
	} else {
		err = EmptyStackError
	}
	return
}

func (m *MyStack) Top() (interface{}, error) {
	if len(m.arr) > 0 {
		return m.arr[len(m.arr)-1], nil
	}
	return nil, EmptyStackError
}

func (m *MyStack) Size() interface{} {
	return len(m.arr)
}
