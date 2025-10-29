package problems

import "fmt"

func summaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}
	var res []string
	start := 0

	for i := 1; i < len(nums); i++ {
		// If current number is NOT consecutive
		if nums[i] != nums[i-1]+1 {
			// Close the current range
			if start == i-1 {
				res = append(res, fmt.Sprintf("%d", nums[start]))
			} else {
				res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[i-1]))
			}
			start = i // Start new range
		}
	}

	// Handle last range
	if start == len(nums)-1 {
		res = append(res, fmt.Sprintf("%d", nums[start]))
	} else {
		res = append(res, fmt.Sprintf("%d->%d", nums[start], nums[len(nums)-1]))
	}

	return res
}
