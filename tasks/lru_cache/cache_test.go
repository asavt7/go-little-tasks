package lru_test

import (
	lru "github.com/asavt7/go-little-tasks/tasks/lru_cache"
	"strconv"
	"testing"
)

func TestLruCache_CalcWithCache(t *testing.T) {

	t.Run("cache work - no unnecessary calls", func(t *testing.T) {
		size := 10
		callCounter := 0
		cache := lru.NewLruCache(size, func(key string) interface{} {
			callCounter++
			return key
		})
		for i := 0; i < size*100; i++ {
			cache.CalcWithCache(strconv.Itoa(i % 10))
		}
		if callCounter != size {
			t.Error("err")
		}
	})
}
