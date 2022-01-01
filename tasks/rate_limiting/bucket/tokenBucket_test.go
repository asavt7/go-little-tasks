package bucket_test

import (
	"github.com/asavt7/go-little-tasks/tasks/rate_limiting/bucket"
	"sync"
	"testing"
	"time"
)

func TestBucket_AllowRequest(t *testing.T) {
	b := bucket.NewBucket(10, 10)

	t.Run("simple test", func(t *testing.T) {

		wg := sync.WaitGroup{}
		wg.Add(10)
		for i := 0; i < 10; i++ {
			go func() {
				if !b.AllowRequest() {
					t.Error("err")
				}
				wg.Done()
			}()
		}
		wg.Wait()

		// empty bucket
		b.AllowRequest()
		b.AllowRequest()

		if b.AllowRequest() {
			t.Error("should return false")
		}

		// refill bucket
		time.Sleep(300 * time.Millisecond)

		if !b.AllowRequest() {
			t.Error("should allow req")
		}
	})

}
