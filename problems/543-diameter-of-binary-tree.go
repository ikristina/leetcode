package problems

func diameterOfBinaryTree(root *TreeNode) int {
	totalMaxDepth := 0

	if root == nil {
		return 0
	}

	var doBT func(node *TreeNode) int

	doBT = func(node *TreeNode) int {
		if node == nil {
			return 0
		}

		rightDepth := doBT(node.Right)
		leftDepth := doBT(node.Left)

		totalDepth := rightDepth + leftDepth
		if totalMaxDepth < totalDepth {
			totalMaxDepth = totalDepth
		}

		if rightDepth > leftDepth {
			return rightDepth + 1
		}

		return leftDepth + 1

	}

	doBT(root)
	return totalMaxDepth
}
