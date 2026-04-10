package problems

func connect(root *TreeNode) *TreeNode {
	if root == nil {
        return nil
    }
    // LInked List
    // leftmost := root
    // for leftmost.Left != nil {

    //     head := leftmost

    //     for head != nil {
    //         head.Left.Next = head.Right
    //         if head.Next != nil {
    //             head.Right.Next = head.Next.Left
    //         }
    //         head = head.Next
    //     }
    //     leftmost = leftmost.Left
    // }

    // BFS

    q := []*TreeNode{root}
    for len(q) > 0 {
        qlen := len(q)

        for i := 0; i < qlen; i++ {
            node := q[0]
            q = q[1:]

            if i < qlen-1 {
                node.Next = q[0]
            }

            if node.Left != nil {
                q = append(q, node.Left)
                q = append(q, node.Right)
            }
        }
    }

    return root
}
