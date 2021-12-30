package lru_test

import (
	lru "github.com/asavt7/go-little-tasks/tasks/lru_cache"
	lru2 "github.com/asavt7/go-little-tasks/tasks/lru_cache_shards"
	"strconv"
	"sync"
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

func BenchmarkCacheSimple(b *testing.B) {
	c := lru2.NewLruCache(100, func(key string) interface{} {
		return key + "test"
	})
	b.ResetTimer()
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = c.CalcWithCache(strconv.Itoa(i))
		}(i)
	}
	wg.Wait()
}

func BenchmarkLruShardCache(b *testing.B) {
	c := lru2.NewLruShardCache(10, 100, func(key string) interface{} {
		return key + "test"
	})
	b.ResetTimer()
	wg := sync.WaitGroup{}
	for i := 0; i < b.N; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			_ = c.CalcWithCache(strconv.Itoa(i))
		}(i)
	}
	wg.Wait()
}
