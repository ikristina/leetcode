package problems

import "math"

func findMin(nums []int) int {
    result := math.MaxInt

    left := 0
    right := len(nums) - 1

    for left <= right {
        mid := (left + right)/2

        if nums[left] <= nums[mid] {
            // the left side is sorted
            if result > nums[left] {
                result = nums[left]
            }
            left = mid + 1
        } else {
            // the right half is sorted
            if result > nums[mid] {
                result = nums[mid]
            }
            right = mid - 1
        }
    }
    return result
}
