package problems

func insertInterval(existingIntervals [][]int, newInterval []int) [][]int {
	if len(existingIntervals) == 0 {
		return [][]int{newInterval}
	}

	res := make([][]int, 0)

	i := 0
	for i < len(existingIntervals) && existingIntervals[i][1] < newInterval[0] {
		res = append(res, existingIntervals[i])
		i++
	}

	for i < len(existingIntervals) && existingIntervals[i][0] <= newInterval[1] {
		newInterval[0] = min(newInterval[0], existingIntervals[i][0])
		newInterval[1] = max(newInterval[1], existingIntervals[i][1])
		i++
	}
	res = append(res, newInterval)
	for i < len(existingIntervals) {
		res = append(res, existingIntervals[i])
		i++
	}
	return res
}
