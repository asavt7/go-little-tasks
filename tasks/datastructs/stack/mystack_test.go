package stack_test

import (
	"github.com/asavt7/go-little-tasks/tasks/datastructs/stack"
	"testing"
)

func TestMyStack_Size(t *testing.T) {

	t.Run("size for empty stack == 0", func(t *testing.T) {
		myStack := stack.NewMyStack()
		if myStack.Size() != 0 {
			t.Error("empty stack size != 0")
		}
	})

	t.Run("size for not empty stack", func(t *testing.T) {
		myStack := stack.NewMyStack()
		expectedSize := 10
		for i := 0; i < expectedSize; i++ {
			myStack.Push(i)
		}
		if myStack.Size() != expectedSize {
			t.Errorf("expected size %d, but  actual %d", expectedSize, myStack.Size())
		}
	})

	t.Run("size for not empty stack after pop and push", func(t *testing.T) {
		myStack := stack.NewMyStack()
		expectedSize := 10
		for i := 0; i < expectedSize; i++ {
			myStack.Push(i)
			_, _ = myStack.Pop()
		}
		if myStack.Size() != 0 {
			t.Errorf("expected size 0, but  actual %d", myStack.Size())
		}
	})
}

func TestMyStack_Top(t *testing.T) {

	t.Run("top on empty stack returns error", func(t *testing.T) {
		myStack := stack.NewMyStack()
		top, err := myStack.Top()
		if err != stack.EmptyStackError {
			t.Errorf("Error shoul be %v, but got %v", stack.EmptyStackError, err)
		}
		if top != nil {
			t.Errorf("empty stack top must be nil value")
		}
	})

	t.Run("top returns last pushed value", func(t *testing.T) {
		myStack := stack.NewMyStack()
		lastVal := 10
		for i := 0; i <= lastVal; i++ {
			myStack.Push(i)
		}
		top, err := myStack.Top()
		if err != nil {
			t.Errorf("err should be nil")
		}
		if top != lastVal {
			t.Errorf("Top() should return last pushed value")
		}
	})

	t.Run("top returns last pushed value - many times", func(t *testing.T) {
		myStack := stack.NewMyStack()
		lastVal := 10
		for i := 0; i <= lastVal; i++ {
			myStack.Push(i)
		}
		for i := 0; i < 5; i++ {
			top, err := myStack.Top()
			if err != nil {
				t.Errorf("err should be nil")
			}
			if top != lastVal {
				t.Errorf("Top() should return last pushed value")
			}
		}
	})

}

func TestMyStack_Pop(t *testing.T) {

	t.Run("pop on empty stack returns error", func(t *testing.T) {
		myStack := stack.NewMyStack()
		top, err := myStack.Pop()
		if err != stack.EmptyStackError {
			t.Errorf("Error shoul be %v, but got %v", stack.EmptyStackError, err)
		}
		if top != nil {
			t.Errorf("empty stack pop must be nil value")
		}
	})

	t.Run("pop returns last pushed value", func(t *testing.T) {
		myStack := stack.NewMyStack()
		lastVal := 10
		for i := 0; i <= lastVal; i++ {
			myStack.Push(i)
		}
		top, err := myStack.Top()
		if err != nil {
			t.Errorf("err should be nil")
		}
		if top != lastVal {
			t.Errorf("Pop() should return last pushed value")
		}
	})

	t.Run("top returns last pushed value - many times", func(t *testing.T) {
		myStack := stack.NewMyStack()
		lastVal := 10
		for i := 0; i <= lastVal; i++ {
			myStack.Push(i)
		}
		for i := lastVal; i >= 0; i-- {
			top, err := myStack.Pop()
			if err != nil {
				t.Errorf("err should be nil")
			}
			if top != i {
				t.Errorf("Pop() should return last pushed value expected=%d, actual %d", i, top)
			}
		}
	})
}
