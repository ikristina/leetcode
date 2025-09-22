package problems

func canJump(nums []int) bool {
	if len(nums) <= 1 {
		return true
	}

	jumpInd := 0
	for i := 0; i <= len(nums)-1; i++ {
		/*if i < jumpInd {
		    continue
		}*/
		if i > jumpInd {
			return false
		}

		if i+nums[i] > jumpInd {
			jumpInd = i + nums[i]
		}
	}
	/*
	   ind := nums[len(nums)-1] + len(nums) - 1
	   i := len(nums) - 1

	   for i > 0 {
	       ind = ind - nums[i] - i
	       i--
	       if ind >= i {
	           return false
	       }
	   }

	*/
	return true
}
