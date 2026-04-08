from common_types import TreeNode
from p102_binary_tree_level_order_traversal import Solution


def test_levelOrder():
    sol = Solution()
    root = TreeNode(3)
    root.left = TreeNode(9)
    root.right = TreeNode(20)
    root.right.left = TreeNode(15)
    root.right.right = TreeNode(7)
    assert sol.levelOrder(root) == [[3], [9, 20], [15, 7]]


def test_levelOrder_empty():
    sol = Solution()
    assert sol.levelOrder(None) == []


def test_levelOrder_single():
    sol = Solution()
    root = TreeNode(3)
    assert sol.levelOrder(root) == [[3]]
