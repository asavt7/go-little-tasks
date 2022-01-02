package FixedWindow

import (
	"fmt"
	"sync/atomic"
	"time"
)

type FixedWindowLimiter struct {
	currentRqCount, maxRq int64
	windowSize            time.Duration
}

func NewFixedWindowLimiter(maxRq int, windowSize time.Duration) *FixedWindowLimiter {
	r := &FixedWindowLimiter{maxRq: int64(maxRq), windowSize: windowSize}
	r.startUpdateWindowWorker()
	return r
}

func (r *FixedWindowLimiter) startUpdateWindowWorker() {
	go func() {
		for {
			time.Sleep(r.windowSize)
			atomic.SwapInt64(&(r.currentRqCount), 0)
		}
	}()
}

func (r *FixedWindowLimiter) AllowRequest() bool {
	if r.currentRqCount >= r.maxRq {
		fmt.Println("blocked")
		return false
	}
	atomic.AddInt64(&(r.currentRqCount), 1)
	fmt.Printf("allowd rq_count %d\n", r.currentRqCount)
	return true
}
