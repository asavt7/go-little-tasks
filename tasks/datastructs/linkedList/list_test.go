package linkedList_test

import (
	"github.com/asavt7/go-little-tasks/tasks/datastructs/linkedList"
	"testing"
)

func TestMyLinkedList(t *testing.T) {

	t.Run("empty_list", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		if ll.Size() != 0 {
			t.Error("size of empty list should be 0")
		}
	})

	t.Run("add list", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		count := 5
		for i := 0; i < count; i++ {
			ll.Add(i)
		}
		if ll.Size() != count {
			t.Errorf("size of  list should be %d", count)
		}
		for i := 0; i < count; i++ {
			get, err := ll.Get(i)
			if err != nil {
				t.Errorf("Get() should not returns errors")
			}
			if get != i {
				t.Errorf("Get() expected %d actual %d", i, get)
			}
		}
	})

	t.Run("add and delete list", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		count := 5
		for i := 0; i < count; i++ {
			ll.Add(i)
		}
		for i := count - 1; i >= 0; i-- {
			err := ll.Delete(i)
			if err != nil {
				t.Errorf("Delete() should not return error")
			}
		}
	})

	t.Run("get from empty list", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		_, err := ll.Get(0)
		if err == nil {
			t.Errorf("Delete() should return error")
		}
	})

	t.Run("get from empty list", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		_, err := ll.Get(0)
		if err == nil {
			t.Errorf("Delete() should return error")
		}
	})

	t.Run("delete from empty list", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		err := ll.Delete(0)
		if err == nil {
			t.Errorf("Delete() should return error")
		}
	})

	t.Run("delete - index out of range ", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		ll.Add(1)
		err := ll.Delete(100)
		if err == nil {
			t.Errorf("Delete() should return error")
		}
	})

	t.Run("find - empty", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		find := ll.Find(1)
		if find >= 0 {
			t.Errorf("should return -1 for empty list")
		}
	})

	t.Run("find - not found", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		ll.Add(1)
		ll.Add(2)
		find := ll.Find(3)
		if find >= 0 {
			t.Errorf("should return -1 for empty list")
		}
	})

	t.Run("find - ok", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		ll.Add(1)
		ll.Add(2)
		ll.Add(3)

		find := ll.Find(3)
		if find != 2 {
			t.Errorf("should return 2 index for data=3")
		}
	})

	t.Run("delete and find ", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		ll.Add(1)
		ll.Add(2)
		ll.Add(3)
		ll.Add(4)
		ll.Add(5)

		for i := 0; i < 3; i++ {
			err := ll.Delete(0)
			if err != nil {
				t.Errorf("should not returns error")
			}
		}

		find := ll.Find(5)
		if find != 1 {
			t.Errorf("should return 1 index for data=5")
		}
	})

	t.Run("delete all and size ", func(t *testing.T) {
		ll := linkedList.NewMyLinkedList()
		ll.Add(1)
		ll.Add(2)
		ll.Add(3)
		ll.Add(4)
		ll.Add(5)
		_ = ll.Delete(4)
		_ = ll.Delete(3)
		_ = ll.Delete(2)
		_ = ll.Delete(1)
		_ = ll.Delete(0)

		size := ll.Size()
		if size != 0 {
			t.Errorf("size should be 0")
		}
	})

	t.Run("working with structs", func(t *testing.T) {
		s := struct {
			a string
		}{a: "1"}

		ll := linkedList.NewMyLinkedList()
		ll.Add(s)
		index := ll.Find(struct {
			a string
		}{a: "1"})
		if index != 0 {
			t.Errorf("Find() should return 0")
		}

		_ = ll.Delete(0)

		size := ll.Size()
		if size != 0 {
			t.Errorf("size should be 0")
		}
	})

}
