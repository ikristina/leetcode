package problems

func averageOfLevels(root *TreeNode) []float64 {
    if root == nil {
        return []float64{}
    }
    averages := []float64{}
    q := []*TreeNode{root}

    for len(q) > 0 {
        qlen := len(q)
        var s float64
        for i := 0; i < qlen; i++ {
            node := q[0]
            q = q[1:]

            s += float64(node.Val)

            if node.Left != nil {
                q = append(q, node.Left)
            }
            if node.Right != nil {
                q = append(q, node.Right)
            }

        }
        a := s/float64(qlen)
        averages = append(averages, a)
    }
    return averages
}
