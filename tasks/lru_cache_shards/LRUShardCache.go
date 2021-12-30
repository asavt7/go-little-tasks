package lru

import (
	"hash"
	"hash/fnv"
	"sync"
)

type LruShardCache struct {
	pool sync.Pool

	shards []*LruCache
}

func (l *LruShardCache) CalcWithCache(key string) interface{} {
	shard := l.getShard(key)
	return shard.CalcWithCache(key)
}

func (l *LruShardCache) getShard(key string) Cache {
	h := (l.pool.Get()).(hash.Hash32)
	h.Reset()
	h.Write([]byte(key))
	l.pool.Put(h)
	return l.shards[int(h.Sum32())%len(l.shards)]
}

func NewLruShardCache(sharsNum int, shardCap int, foo func(key string) interface{}) *LruShardCache {
	shards := make([]*LruCache, sharsNum)
	for i := 0; i < sharsNum; i++ {
		shards[i] = NewLruCache(shardCap, foo)
	}

	pool := sync.Pool{New: func() interface{} { return fnv.New32a() }}
	for i := 0; i < sharsNum; i++ {
		pool.Put(fnv.New32a())
	}

	return &LruShardCache{shards: shards, pool: pool}
}
