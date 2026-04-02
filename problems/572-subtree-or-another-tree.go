package problems

func isSubtree(tree *TreeNode, subtree *TreeNode) bool {
	if tree == nil {
		return false
	}

	return isSameTree(tree, subtree) || isSubtree(tree.Left, subtree) || isSubtree(tree.Right, subtree)
}

func isSameTree(a, b *TreeNode) bool {
	if a == nil || b == nil {
		return a == b
	}

	if a.Val != b.Val {
		return false
	}

	return isSameTree(a.Left, b.Left) && isSameTree(a.Right, b.Right)
}
