package problems


type TreeNode struct {
      Val int
      Left *TreeNode
      Right *TreeNode
}

func maxDepth(root *TreeNode) int {
    if root == nil {
        return 0
    }
    // return 1 + max(maxDepth(root.Left), maxDepth(root.Right))

    queue := []*TreeNode{root}
    depth := 0

    for len(queue) > 0 {
        depth ++

        level_size := len(queue)

        for range level_size {
            node := queue[0]
            
            queue = queue[1:]

            if node.Left != nil {
                queue = append(queue, node.Left)
            }
            if node.Right != nil {
                queue = append(queue, node.Right)
            }

        }

    } 
    return depth

}