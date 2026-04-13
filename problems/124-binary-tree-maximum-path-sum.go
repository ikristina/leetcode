package problems

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32

	var dfs func(root *TreeNode) int
	dfs = func(root *TreeNode) int {
		if root == nil {
			return 0
		}

		left_gain := max(dfs(root.Left), 0)
		right_gain := max(dfs(root.Right), 0)

		maxSum = max(maxSum, left_gain + right_gain + root.Val)

		return root.Val + max(left_gain, right_gain)
	}

	dfs(root)
	return maxSum
}
