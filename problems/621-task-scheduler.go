package problems

import "container/heap"

type pair struct {
    count int
    task  byte
}

type schedulerHeap []pair

func (h schedulerHeap) Len() int            { return len(h) }
func (h schedulerHeap) Less(i, j int) bool  { return h[i].count > h[j].count }
func (h schedulerHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }

func (h *schedulerHeap) Push(x any)         { *h = append(*h, x.(pair)) }
func (h *schedulerHeap) Pop() any {
    old := *h
    n := len(old)
    x := old[n-1]
    *h = old[0 : n-1]
    return x
}


func leastInterval(tasks []byte, n int) int {
    freqs := map[byte]int{}
    for _, task := range tasks {
        freqs[task]++
    }

    forH := []pair{}
    for k, v := range freqs {
        p := pair{
            count: v,
            task: k,
        }
        forH = append(forH, p)
    }


    h := schedulerHeap(forH)
    heap.Init(&h)

    cpuCount := 0
    for h.Len() > 0 {
        cycle := n + 1
        used := 0

        temp := []pair{}

        for i := 0; i < cycle; i++ {
            if h.Len() != 0 {
                pair := heap.Pop(&h).(pair)
                temp = append(temp, pair)
                used++
            } else {
                temp = append(temp, pair{count: 0, task: 0})
            }
        }

        for _, t := range temp {
            if t.count - 1 > 0 {
                heap.Push(&h, pair{count:t.count-1, task: t.task})
            }
        }

        if h.Len() > 0 {
            cpuCount = cpuCount + cycle
        } else {
            cpuCount = cpuCount + used
        }

    }

    return cpuCount
}
