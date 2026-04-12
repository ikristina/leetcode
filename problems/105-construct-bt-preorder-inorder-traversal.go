package problems

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func buildTree(preorder []int, inorder []int) *TreeNode {
    mapping := make(map[int]int)
    for i, v := range inorder {
        mapping[v] = i
    }
    pInd := 0

    var buildTreeHelper func(leftInd, rightInd int) *TreeNode

    buildTreeHelper = func (leftInd, rightInd int) *TreeNode {
        if leftInd > rightInd {
            return nil
        }
        val := preorder[pInd]
        pInd++
        root := &TreeNode{Val: val}

        iInd := mapping[val]

        root.Left = buildTreeHelper(leftInd, iInd-1)
        root.Right = buildTreeHelper(iInd+1, rightInd)

        return root
    }

    return buildTreeHelper(0, len(preorder)-1)
}
