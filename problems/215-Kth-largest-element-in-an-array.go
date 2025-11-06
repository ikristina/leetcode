package problems

import (
	"container/heap"
	"fmt"
)

type maxHeap []int

func (h maxHeap) Len() int {
	return len(h)
}

func (h maxHeap) Less(i, j int) bool {
	return h[i] > h[j]
}

func (h maxHeap) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

func (h *maxHeap) Push(i any) {
	*h = append(*h, i.(int))
}

func (h *maxHeap) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func findKthLargest(nums []int, k int) int {
	h := maxHeap(nums)
	heap.Init(&h)

	x := 0
	for range k {
		x = heap.Pop(&h).(int)
		fmt.Println(x)
	}

	return x
}
