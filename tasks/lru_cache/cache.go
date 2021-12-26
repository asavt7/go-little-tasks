package lru

import (
	"container/heap"
	"time"
)

type CachedItem struct {
	key       string
	value     interface{}
	timestamp time.Time
	index     int
}

type LruCache struct {
	maxSize       int
	cache         map[string]*CachedItem
	queue         *Queue
	calculateFunc func(key string) interface{}
}

func NewLruCache(maxSize int, calculateFunc func(key string) interface{}) *LruCache {
	q := make(Queue, 0, maxSize)
	return &LruCache{
		maxSize:       maxSize,
		cache:         make(map[string]*CachedItem),
		queue:         &q,
		calculateFunc: calculateFunc,
	}
}

func (c *LruCache) CalcWithCache(key string) interface{} {
	now := time.Now()

	if v, ok := c.cache[key]; ok {
		c.queue.update(v, now)
		return v
	}

	value := c.calculateFunc(key)

	if len(c.cache) == c.maxSize {
		item := heap.Pop(c.queue).(*CachedItem)
		keyToRemove := item.key
		delete(c.cache, keyToRemove)
	}

	v := &CachedItem{
		key:       key,
		value:     value,
		timestamp: now,
	}
	c.cache[key] = v
	heap.Push(c.queue, v)

	return value
}
