package problems

func rotate(nums []int, k int) {
	if k == 0 {
		return
	}
	if len(nums) == 0 {
		return
	}
	k %= len(nums) // to stay within the length of nums.

	reverse(nums, 0, len(nums)-1) // reverse the whole slice so the numbers are in the correct order
	reverse(nums, 0, k-1)         // reverse first part to restore the order
	reverse(nums, k, len(nums)-1)

}

func reverse(nums []int, p1, p2 int) {
	for p1 < p2 {
		nums[p1], nums[p2] = nums[p2], nums[p1]
		p1++
		p2--
	}
}
