package problems

import "container/heap"

type DistMaxHeap [][]int // []int{distance, index}

func (h DistMaxHeap) Len() int { return len(h) }
func (h DistMaxHeap) Less(i, j int) bool {return h[i][0] > h[j][0]}
func (h DistMaxHeap) Swap(i, j int) {h[i], h[j] = h[j], h[i]}

func (h *DistMaxHeap) Push(x any) {
    // Push and Pop use pointer receivers because they modify the slice's length,
	// not just its contents.
	*h = append(*h, x.([]int))
}

func (h *DistMaxHeap) Pop() any {
    old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}

func kClosest(points [][]int, k int) [][]int {
    h := DistMaxHeap([][]int{})
    heap.Init(&h)
    result := [][]int{}
    for i := 0; i < len(points); i++ {
        x, y := points[i][0], points[i][1]
        dist := x*x + y*y

        if h.Len() == k {
            root := heap.Pop(&h).([]int)
            if root[0] >= dist {
                heap.Push(&h, []int{dist, i})
            } else {
                heap.Push(&h, root)
            }
            continue
        }
        heap.Push(&h, []int{dist, i})
    }

    for i := 0; i < k; i++ {
        item := heap.Pop(&h).([]int)
        result = append(result, points[item[1]])
    }

    return result
}
