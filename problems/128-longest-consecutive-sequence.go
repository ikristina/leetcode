package problems

func longestConsecutive(nums []int) int {
	hm := map[int]struct{}{}
	for _, n := range nums {
		hm[n] = struct{}{}
	}

	ml := 0 // max length
	for k := range hm {
		if _, ok := hm[k-1]; ok {
			// not the first element in sequence
			continue
		}

		l := 1
		for {
			if _, ok := hm[k+l]; ok {
				l++
				continue
			}
			break
		}

		if l > ml {
			ml = l
		}
	}

	return ml
}
