package SlicingWindow_test

import (
	"github.com/asavt7/go-little-tasks/tasks/rate_limiting/SlicingWindow"
	"sync"
	"testing"
	"time"
)

func TestFixedWindowLimiter_AllowRequest(t *testing.T) {
	r := SlicingWindow.NewSlicingWindowLimiter(10, time.Second)

	for i := 0; i < 10; i++ {
		if !r.AllowRequest() {
			t.Error("should allow rq")
		}
	}
	time.Sleep(1000 * time.Millisecond)
	time.Sleep(50 * time.Millisecond)
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()
	r.AllowRequest()

	if r.AllowRequest() {
		t.Error("should not allow rq")
	}

}

func TestFixedWindowLimiter_AllowRequest_concurrent(t *testing.T) {
	r := SlicingWindow.NewSlicingWindowLimiter(10, time.Second)

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
	time.Sleep(1000 * time.Millisecond)
	r.AllowRequest()
	if !r.AllowRequest() {
		t.Error("should not allow rq")
	}

	time.Sleep(300 * time.Millisecond)
	if !r.AllowRequest() {
		t.Error("should not allow rq")
	}
}
