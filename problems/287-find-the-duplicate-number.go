package problems

// FindDuplicate solves LeetCode 287: Find the Duplicate Number
// Time Complexity: O(n)
// Space Complexity: O(1)
// Fast and slow pointer
func FindDuplicate(nums []int) int {
	sp := nums[0]
	fp := nums[0]

	for {
		sp = nums[sp]
		fp = nums[nums[fp]]
		if sp == fp {
			break
		}
	}

	sp = nums[0]
	for sp != fp {
		sp = nums[sp]
		fp = nums[fp]
	}
	return sp
}
