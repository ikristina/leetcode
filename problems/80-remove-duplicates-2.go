package problems

func removeDuplicates2(nums []int) int {
	p1 := 0
	p2 := p1 + 1

	howMany := 0
	for p2 < len(nums) {
		if nums[p1] == nums[p2] {
			howMany++
			if howMany == 2 {
				nums = append(nums[:p1], nums[p2:]...)
				howMany--
				continue
			}
			p1 = p2
			p2++
		} else {
			howMany = 0
			p1 = p2
			p2++
		}
	}
	return len(nums)
}
