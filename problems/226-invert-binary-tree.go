package problems

// Definition of a binary tree node
// type TreeNode[T any] struct {
// 	Data  T
// 	Left  *TreeNode[T]
// 	Right *TreeNode[T]
// }

func mirrorBinaryTree(root *TreeNode) *TreeNode {

	if root == nil {
		return root
	}

	queue := []*TreeNode{root}

	for len(queue) > 0 {

		level_size := len(queue)

		for level_size > 0 {

			node := queue[0]
			queue = queue[1:]

			node.Left, node.Right = node.Right, node.Left

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			level_size--
		}
	}
	// root.Left, root.Right = mirrorBinaryTree(root.Right), mirrorBinaryTree(root.Left)
	return root
}
