package SlicingWindow

import (
	"fmt"
	"sync"
	"time"
)

type SlicingWindowLimiter struct {
	mu sync.Mutex

	currentRqCount, previousWindowRqCount, maxRq int64
	windowSize                                   time.Duration
	windowStarted                                time.Time
}

func NewSlicingWindowLimiter(maxRq int, windowSize time.Duration) *SlicingWindowLimiter {
	r := &SlicingWindowLimiter{
		mu:                    sync.Mutex{},
		currentRqCount:        0,
		previousWindowRqCount: 0,
		maxRq:                 int64(maxRq),
		windowSize:            windowSize,
		windowStarted:         time.Now(),
	}

	r.startUpdateWindowWorker()
	return r
}

func (r *SlicingWindowLimiter) startUpdateWindowWorker() {
	go func() {
		for {
			time.Sleep(r.windowSize)
			r.mu.Lock()
			r.previousWindowRqCount = r.currentRqCount
			r.windowStarted = time.Now()
			r.currentRqCount = 0
			r.mu.Unlock()
		}
	}()
}

func (r *SlicingWindowLimiter) AllowRequest() bool {
	r.mu.Lock()
	defer r.mu.Unlock()

	part := float64(time.Now().UnixNano()-r.windowStarted.UnixNano()) / float64(r.windowSize.Nanoseconds())

	slicingCount := float64(r.currentRqCount)*part + float64(r.previousWindowRqCount)*(1-part)
	fmt.Printf("slicing count %f\n", slicingCount)
	if slicingCount >= float64(r.maxRq) {
		fmt.Println("blocked")
		return false
	}
	r.currentRqCount++
	fmt.Printf("allowd rq_count %d\n", r.currentRqCount)
	return true
}
