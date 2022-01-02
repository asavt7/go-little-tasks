package FixedWindow_test

import (
	"github.com/asavt7/go-little-tasks/tasks/rate_limiting/FixedWindow"
	"sync"
	"testing"
	"time"
)

func TestFixedWindowLimiter_AllowRequest(t *testing.T) {
	r := FixedWindow.NewFixedWindowLimiter(10, time.Second)

	for i := 0; i < 10; i++ {
		if !r.AllowRequest() {
			t.Error("should allow rq")
		}
	}

	if r.AllowRequest() {
		t.Error("should not allow rq")
	}

}

func TestFixedWindowLimiter_AllowRequest_concurrent(t *testing.T) {
	r := FixedWindow.NewFixedWindowLimiter(10, time.Second)

	wg := sync.WaitGroup{}
	wg.Add(10)
	for i := 0; i < 10; i++ {
		go func() {
			if !r.AllowRequest() {
				t.Error("should allow rq")
			}
			wg.Done()
		}()
	}
	wg.Wait()
	if r.AllowRequest() {
		t.Error("should not allow rq")
	}
	time.Sleep(1100 * time.Millisecond)
	if !r.AllowRequest() {
		t.Error("should  allow rq")
	}

}
