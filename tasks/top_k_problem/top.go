package main

import (
	"container/heap"
	"fmt"
)

func topK(nums []int, k int) []int {
	m := make(map[int]int)

	for _, num := range nums {
		m[num]++
	}

	h := make(IntHeap, 0, k)
	for key := range m {
		fmt.Printf("%d = %d\n", key, m[key])

		if len(h) == k {
			x := heap.Pop(&h)
			if x.(Count).count > m[key] {
				heap.Push(&h, x)
				continue
			}
		}
		heap.Push(&h, Count{
			key:   key,
			count: m[key],
		})
	}

	res := make([]int, 0, k)
	for len(h) > 0 {
		res = append(res, heap.Pop(&h).(Count).key)
	}
	return res
}

type Count struct {
	key   int
	count int
}

type IntHeap []Count

func (r *IntHeap) Len() int {
	return len(*r)
}

func (r *IntHeap) Less(i, j int) bool {
	return (*r)[i].count < (*r)[j].count
}

func (r *IntHeap) Swap(i, j int) {
	(*r)[i], (*r)[j] = (*r)[j], (*r)[i]
}

func (r *IntHeap) Push(x interface{}) {
	*r = append(*r, x.(Count))
}

func (r *IntHeap) Pop() interface{} {
	old := *r
	n := len(old)
	x := old[n-1]
	*r = old[0 : n-1]
	return x
}

func main() {
	a := []int{1, 1, 1, 1, 1, 1, 1, 1, 1, 2, 3, 3, 3, 3, 2, 2, 2, 2, 2, 2, 2, 2, 2, 6, 6, 6, 6, 6, 6, 7, 7, 8, 8, 8, 9, 0, 11, 12, 33, 5, 6, 7}
	fmt.Println(topK(a, 3))
}
