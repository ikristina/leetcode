package problems

func spiralOrder(matrix [][]int) []int {

	if len(matrix) == 0 {
		return []int{}
	}
	res := []int{}

	top := 0
	bottom := len(matrix) - 1
	left := 0
	right := len(matrix[0]) - 1

	for len(res) < len(matrix)*len(matrix[0]) {
		if top <= bottom {
			for i := left; i <= right; i++ {
				res = append(res, matrix[top][i])
			}
			top++
		}
		if left <= right {
			for i := top; i <= bottom; i++ {
				res = append(res, matrix[i][right])
			}
			right--
		}

		if top <= bottom {
			for i := right; i >= left; i-- {
				res = append(res, matrix[bottom][i])
			}
			bottom--
		}

		if left <= right {
			for i := bottom; i >= top; i-- {
				res = append(res, matrix[i][left])
			}
			left++
		}

	}
	return res
}
