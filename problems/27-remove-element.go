package problems

func RemoveElement(nums []int, val int) int {
	lp := 0
	rp := len(nums) - 1
	for len(nums) > 0 && lp < len(nums) {
		if nums[lp] == val {
			nums = append(nums[:lp], nums[lp+1:]...)
			rp--
		} else {
			lp++
		}

	}
	return len(nums)
}
