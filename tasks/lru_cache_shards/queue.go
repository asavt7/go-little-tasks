package lru

import (
	"container/heap"
	"time"
)

type Queue []*CachedItem

func (q Queue) Len() int {
	return len(q)
}

func (q Queue) Less(i, j int) bool {
	return q[i].timestamp.UnixNano() < q[j].timestamp.UnixNano()
}

func (q Queue) Swap(i, j int) {
	q[i], q[j] = q[j], q[i]
	q[i].index = i
	q[j].index = j
}

func (q *Queue) Push(x interface{}) {
	n := len(*q)
	item := x.(*CachedItem)
	item.index = n
	*q = append(*q, item)
}

func (q *Queue) Pop() interface{} {
	old := *q
	n := len(old)
	item := old[n-1]
	old[n-1] = nil  // avoid memory leak
	item.index = -1 // for safety
	*q = old[0 : n-1]
	return item
}

func (q *Queue) update(item *CachedItem, lastUsed time.Time) {
	item.timestamp = lastUsed
	heap.Fix(q, item.index)
}
