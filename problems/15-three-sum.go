package problems

import "sort"

// ThreeSum https://leetcode.com/problems/3sum
// medium
// #array, #two pointers
func ThreeSum(nums []int) [][]int {

	sort.Ints(nums)

	r := [][]int{}

	for i := 0; i < len(nums)-2; i++ {

		if nums[i] > 0 {
			break
		}

		if i == 0 || nums[i] != nums[i-1] {

			low, high := i+1, len(nums)-1

			for low < high {
				s := nums[i] + nums[low] + nums[high]

				if s > 0 {
					high--
				} else if s < 0 {
					low++
				} else {
					r = append(r, []int{nums[i], nums[low], nums[high]})
					low++
					high--

					for low < high && nums[low] == nums[low-1] {
						low++
					}

					for low < high && nums[high] == nums[high+1] {
						high--
					}

				}

			}
		}

	}

	return r

}
