package problems

/*
Monotonic stack - same pattern as Daily Temperatures
Width = right - left - 1 where left comes from the new stack top after popping

Same pattern as 42 - Trapping Rain Water.
*/

func largestRectangleArea(heights []int) int {
	heights = append(heights, 0)
	lra := 0

	stack := []int{} // indices

	for i, h := range heights {

		for len(stack) > 0 && h < heights[stack[len(stack)-1]] {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]

			left_boundary := -1
			if len(stack) > 0 {
				left_boundary = stack[len(stack)-1]
			}
			width := i - left_boundary - 1
			lra = max(lra, heights[top]*width)
		}

		stack = append(stack, i)

	}

	return lra
}
