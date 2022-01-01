package bucket

import (
	"sync"
	"time"
)

type Bucket struct {
	mu sync.Mutex

	curSize, maxSize    int
	refillRatePerSecond int

	lastUpdate time.Time
}

func NewBucket(maxSize int, refillRatePerSecond int) *Bucket {
	return &Bucket{
		curSize:             maxSize,
		maxSize:             maxSize,
		refillRatePerSecond: refillRatePerSecond,
		lastUpdate:          time.Now(),
	}
}

func (b *Bucket) AllowRequest() bool {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.refill()
	if b.curSize > 0 {
		b.curSize--
		return true
	}
	return false
}

func (b *Bucket) refill() {
	now := time.Now()

	refillCount := int(now.UnixNano()-b.lastUpdate.UnixNano()) * b.refillRatePerSecond / 1e9
	b.lastUpdate = now
	b.curSize += refillCount
	if b.curSize > b.maxSize {
		b.curSize = b.maxSize
	}
}
