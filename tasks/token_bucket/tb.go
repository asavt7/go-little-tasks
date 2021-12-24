package main

import "time"

type TokenBucket struct {
	maxSize          int
	refillRatePerSec int

	currentSize           int
	lastRefillTimestampNs int64
}

func NewTokenBucket(maxSize int, refillRate int) *TokenBucket {
	return &TokenBucket{maxSize: maxSize, refillRatePerSec: refillRate, currentSize: maxSize, lastRefillTimestampNs: time.Now().UnixNano()}
}

func (r *TokenBucket) allowRequest(tokens int) bool {
	r.refill()
	if r.currentSize > tokens {
		r.currentSize -= tokens
		return true
	}
	return false
}

func (r *TokenBucket) refill() {
	now := time.Now().UnixNano()
	toAddTokens := int((now - r.lastRefillTimestampNs) * int64(r.refillRatePerSec) / 1e9)
	curSize := r.currentSize + toAddTokens
	if curSize > r.maxSize {
		r.currentSize = r.maxSize
	} else {
		r.currentSize = curSize
	}
	r.lastRefillTimestampNs = now
}

func main() {
}
