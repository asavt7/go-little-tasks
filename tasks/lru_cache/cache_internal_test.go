package lru

import (
	"container/heap"
	"strconv"
	"testing"
)

func TestLruCache_CalcWithCache(t *testing.T) {

	t.Run("cache 10 keys, each key calls 10 times", func(t *testing.T) {
		size := 10
		callCounter := 0
		cache := NewLruCache(size, func(key string) interface{} {
			callCounter++
			return key
		})
		for i := 0; i < size*10; i++ {
			cache.CalcWithCache(strconv.Itoa(i % 10))
		}
		if callCounter != size {
			t.Error("err")
		}

		if cache.queue.Len() != size {
			t.Error("priority queue size != max size of cache")
		}
		if len(cache.cache) != size {
			t.Error("cache map size != max size of cache")
		}
	})

	t.Run("check queue priority", func(t *testing.T) {
		size := 10
		callCounter := 0
		cache := NewLruCache(size, func(key string) interface{} {
			callCounter++
			return key
		})
		for i := 0; i < size; i++ {
			cache.CalcWithCache(strconv.Itoa(i))
		}

		if v := heap.Pop(cache.queue).(*CachedItem).key; v != "0" {
			t.Errorf("priority queue work with err! expected head: %s, but got %s", "0", v)
		}
	})

	t.Run("check queue priority 1", func(t *testing.T) {
		size := 10
		callCounter := 0
		cache := NewLruCache(size, func(key string) interface{} {
			callCounter++
			return key
		})
		for i := 0; i < size+1; i++ {
			cache.CalcWithCache(strconv.Itoa(i))
		}

		if v := heap.Pop(cache.queue).(*CachedItem).key; v != "1" {
			t.Errorf("priority queue work with err! expected head: %s, but got %s", "1", v)
		}
	})

}
