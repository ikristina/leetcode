package problems

func search(nums []int, target int) int {
	low := 0
	high := len(nums) - 1

	for low <= high {
		mid := (low + high) / 2

		if nums[low] <= nums[mid] {
			if target >= nums[low] && target < nums[mid] {
				if target == nums[low] {
					return low
				}
				high = mid - 1
			} else {
				if target == nums[mid] {
					return mid
				}
				if target == nums[high] {
					return high
				}
				low = mid + 1
			}
		} else {
			if target >= nums[mid] && target <= nums[high] {
				if target == nums[high] {
					return high
				}
				if target == nums[mid] {
					return mid
				}
				low = mid + 1
			} else {
				if target == nums[low] {
					return low
				}
				if target == nums[mid] {
					return mid
				}
				high = mid - 1
			}
		}

	}
	return -1
}
