package problems

func dailyTemperatures(temperatures []int) []int {
	result := make([]int, len(temperatures))
	stack := []int{} // store indices

	for i, t := range temperatures {

		for len(stack) > 0 && t > temperatures[stack[len(stack)-1]] {
			last := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			result[last] = i - last
		}

		stack = append(stack, i)
	}

	return result
}
