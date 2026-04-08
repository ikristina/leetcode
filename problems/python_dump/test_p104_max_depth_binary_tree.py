from p104_max_depth_binary_tree import Solution, TreeNode


def test_max_depth_empty_tree():
    sol = Solution()
    assert sol.maxDepth(None) == 0


def test_max_depth_single_node():
    sol = Solution()
    root = TreeNode(1)
    assert sol.maxDepth(root) == 1


def test_max_depth_complex_tree():
    #      3
    #     / \
    #    9  20
    #      /  \
    #     15   7
    sol = Solution()
    root = TreeNode(3)
    root.left = TreeNode(9)
    root.right = TreeNode(20, TreeNode(15), TreeNode(7))

    assert sol.maxDepth(root) == 3
