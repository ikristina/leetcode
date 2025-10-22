package problems

func minSubArrayLen(target int, nums []int) int {
	resLen := len(nums) + 1 // impossible value

	currSum := 0
	l, r := 0, 0

	for r < len(nums) {

		currSum += nums[r]
		r++

		for currSum >= target && l < r {
			resLen = min(resLen, r-l)
			currSum -= nums[l]
			l++
		}

	}

	if resLen == len(nums)+1 {
		return 0
	}

	return resLen
}
