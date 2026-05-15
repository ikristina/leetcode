package problems

import (
	"container/heap"
	"math"

	"github.com/ikristina/leetcode/data_structures"
)

// networkDelayTimePQ finds the time it takes for all nodes to receive a signal.
// It uses Dijkstra's algorithm with a Priority Queue for O(E log V) efficiency.
func networkDelayTimePQ(times [][]int, n, k int) int {
	// Initialize distances to all nodes as infinity.
	dist := make([]int, n+1)
	for i := range dist {
		dist[i] = math.MaxInt32
	}
	dist[k] = 0

	// Build the adjacency list.
	graph := make(map[int][][]int)
	for _, t := range times {
		graph[t[0]] = append(graph[t[0]], []int{t[1], t[2]})
	}

	// Initialize the Min-Heap with the starting node.
	pq := &data_structures.PriorityQueue{}
	heap.Init(pq)
	heap.Push(pq, data_structures.PriorityQueueItem{Value: k, Priority: 0})

	for pq.Len() > 0 {
		// Pop the node with the SMALLEST distance.
		curr := heap.Pop(pq).(data_structures.PriorityQueueItem)
		u := curr.Value.(int)
		d := curr.Priority

		// Optimization: if we already found a better path to 'u',
		// ignore this stale entry from the heap.
		if d > dist[u] {
			continue
		}

		// Explore neighbors and perform "Relaxation".
		for _, edge := range graph[u] {
			v, weight := edge[0], edge[1]
			if dist[u]+weight < dist[v] {
				dist[v] = dist[u] + weight
				heap.Push(pq, data_structures.PriorityQueueItem{Value: v, Priority: dist[v]})
			}
		}
	}

	// Find the maximum distance to any reachable node.
	maxTime := 0
	for i := 1; i <= n; i++ {
		if dist[i] == math.MaxInt32 {
			return -1 // A node is unreachable.
		}
		if dist[i] > maxTime {
			maxTime = dist[i]
		}
	}
	return maxTime
}
