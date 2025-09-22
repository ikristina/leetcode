package problems

func removeDuplicates(nums []int) int {
	// the slice is sorted.

	p1 := 0
	p2 := p1 + 1

	for p2 <= len(nums) {
		if p2 == len(nums) {
			nums = append(nums[:p1+1], nums[p2:]...)
			break
		}
		if nums[p1] == nums[p2] {
			p2++
		} else if nums[p1] != nums[p2] {
			nums = append(nums[:p1+1], nums[p2:]...)
			p1++
			p2 = p1 + 1
		}
	}
	return len(nums)
}
