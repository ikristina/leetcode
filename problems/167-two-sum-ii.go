package problems

func twoSum2(numbers []int, target int) []int {
	// ascending order

	p1 := 0
	p2 := len(numbers) - 1

	for {
		if p1 >= p2 {
			break
		}
		sum := numbers[p1] + numbers[p2]
		if sum == target {
			return []int{p1 + 1, p2 + 1}
		}

		if sum > target {
			p2--
		}

		if sum < target {
			p1++
		}

	}
	return []int{}
}
