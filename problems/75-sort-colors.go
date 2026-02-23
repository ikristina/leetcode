package problems

// SortColors https://leetcode.com/problems/sort-colors
// medium
func SortColors(nums []int) {
	low, med, high := 0, 0, len(nums)-1

	for med <= high {
		switch nums[med] {
		case 0:
			nums[low], nums[med] = nums[med], nums[low]
			low++
			med++
		case 1:
			med++
		case 2:
			nums[high], nums[med] = nums[med], nums[high]
			high--
		}

	}
}
