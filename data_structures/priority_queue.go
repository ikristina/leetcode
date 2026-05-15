package data_structures

// PriorityQueueItem represents an item in the priority queue.
type PriorityQueueItem struct {
	Value    any
	Priority int
}

// PriorityQueue implements heap.Interface and holds PriorityQueueItems.
// It is a Min-Heap based on the Priority field.
type PriorityQueue []PriorityQueueItem

func (pq PriorityQueue) Len() int           { return len(pq) }
func (pq PriorityQueue) Less(i, j int) bool { return pq[i].Priority < pq[j].Priority }
func (pq PriorityQueue) Swap(i, j int)      { pq[i], pq[j] = pq[j], pq[i] }

func (pq *PriorityQueue) Push(x any) {
	*pq = append(*pq, x.(PriorityQueueItem))
}

func (pq *PriorityQueue) Pop() any {
	old := *pq
	n := len(old)
	item := old[n-1]
	*pq = old[0 : n-1]
	return item
}
