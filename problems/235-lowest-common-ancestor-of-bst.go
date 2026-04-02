package problems

func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil {
		return nil
	}

	if root == p || root == q {
		return root
	}

	// check if they are in the same subtree
	if root.Val > p.Val && root.Val > q.Val {
		root = root.Left
	} else if root.Val < p.Val && root.Val < q.Val {
		root = root.Right
	} else {
		return root
	}

	return lowestCommonAncestor(root, p, q)
	// for root != nil {
	//     if root.Val > p.Val && root.Val > q.Val {
	//         root = root.Left
	//     } else if root.Val < p.Val && root.Val < q.Val {
	//         root = root.Right
	//     } else {
	//         return root
	//     }
	// }
	// return nil
}
