package linkedList

import "errors"

var (
	IndexOfOutOfRangeErr = errors.New("index out of range")
)

type List interface {
	Add(interface{})
	Size() int

	Get(int) (interface{}, error)
	Delete(int) error
	Find(interface{}) int
}

type node struct {
	data interface{}
	next *node
}

type MyLinkedList struct {
	size int
	head *node
}

func NewMyLinkedList() *MyLinkedList {
	return &MyLinkedList{}
}

func (m *MyLinkedList) Add(i interface{}) {
	newNode := &node{
		data: i,
		next: nil,
	}
	defer func() { m.size++ }()

	var current *node = m.head
	if m.size == 0 {
		m.head = newNode
		return
	}
	for current.next != nil {
		current = current.next
	}
	current.next = newNode
}

func (m *MyLinkedList) Size() int {
	return m.size
}

func (m *MyLinkedList) Get(i int) (interface{}, error) {
	if i >= m.size {
		return nil, IndexOfOutOfRangeErr
	}
	var current *node = m.head
	for ; i > 0; i-- {
		current = current.next
	}
	return current.data, nil
}

func (m *MyLinkedList) Delete(i int) error {
	if i >= m.size {
		return IndexOfOutOfRangeErr
	}
	defer func() { m.size-- }()

	if i == 0 {
		m.head = m.head.next
		return nil
	}

	var prev *node = m.head.next
	for indx := i; indx > 1; indx-- {
		prev = prev.next
	}

	if i == (m.size - 1) {
		prev.next = nil
	} else {
		prev.next = prev.next.next
	}
	return nil
}

func (m *MyLinkedList) Find(i interface{}) int {
	var current *node = m.head
	for indx := 0; indx < m.size; indx++ {
		if i == current.data {
			return indx
		}
		current = current.next
	}
	return -1
}
