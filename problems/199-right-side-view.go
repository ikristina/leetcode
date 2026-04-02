package problems

func rightSideView(root *TreeNode) []int {
	res := []int{}

	if root == nil {
		return res
	}

	var dfs func(node *TreeNode, depth int)

	dfs = func(node *TreeNode, depth int) {
		if node == nil {
			return
		}

		if depth == len(res) {
			// first time at this depth
			res = append(res, node.Val)
		}

		depth++

		dfs(node.Right, depth)
		dfs(node.Left, depth)

	}

	dfs(root, 0)
	return res
}
