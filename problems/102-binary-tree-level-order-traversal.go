package problems

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	out := [][]int{}

	queue := []*TreeNode{root}

	for len(queue) > 0 {
		levelLen := len(queue)
		level := []int{}
		for _ = range levelLen {
			node := queue[0]
			queue = queue[1:]

			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}

			level = append(level, node.Val)
		}
		out = append(out, level)
	}

	return out
}
