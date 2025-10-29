package problems

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	res := make([][]int, 0)

	if len(intervals) == 0 {
		return res
	}
	// sort by first element
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})

	start := intervals[0][0]
	end := intervals[0][len(intervals[0])-1]

	for i := 1; i < len(intervals); i++ {
		if len(intervals[i]) == 0 {
			continue
		}
		if intervals[i][0] <= end {
			end = max(end, intervals[i][len(intervals[i])-1])
		} else {
			res = append(res, []int{start, end})
			start = intervals[i][0]
			end = intervals[i][len(intervals[i])-1]
		}
	}

	res = append(res, []int{start, end})

	return res
}
