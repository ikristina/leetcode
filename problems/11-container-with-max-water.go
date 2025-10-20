package problems

func maxArea(height []int) int {
	if len(height) == 0 || len(height) == 1 {
		return 0
	}

	if len(height) == 2 {
		if height[0] > height[1] {
			return height[1] * height[1]
		}
		return height[0] * height[0]
	}

	p1 := 0
	p2 := len(height) - 1

	prod := 1

	for p1 < p2 {

		if height[p1] <= height[p2] {
			if prod < height[p1]*(p2-p1) {
				prod = height[p1] * (p2 - p1)
			}
			p1++

		}

		if height[p1] >= height[p2] {
			if prod < height[p2]*(p2-p1) {
				prod = height[p2] * (p2 - p1)
			}
			p2--
		}

	}
	return prod
}
