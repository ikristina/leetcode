package problems

type HitCounter struct {
	hits []int // queue
}

func HitCounterConstructor() HitCounter {
	return HitCounter{
		hits: make([]int, 0),
	}
}

func (h *HitCounter) Hit(timestamp int) {
	// O(1) amortized append
	h.hits = append(h.hits, timestamp)
}

func (h *HitCounter) GetHits(timestamp int) int {
	if len(h.hits) == 0 {
		return 0
	}

	// binary search
	left, right := 0, len(h.hits)-1
	for left <= right {
		mid := left + (right - left) / 2 // to prevent integer overflow if they are both large numbers. (instead of (left + right) / 2)
		if h.hits[mid] <= timestamp - 300 {
			// too old, look in the right half
			left = mid + 1
		} else {
			// valid but try to find earlier valid in the left half
			right = mid - 1
		}
	}
	// 'left' is now the index of the first valid hit
	arr := make([]int, len(h.hits)-left)
	copy(arr, h.hits[left:])
	h.hits = arr
	return len(h.hits)
}
