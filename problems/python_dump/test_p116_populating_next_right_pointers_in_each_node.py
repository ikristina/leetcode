from p116_populating_next_right_pointers_in_each_node import Node, Solution


def test_connect():
    solution = Solution()
    root = Node(1, Node(2, Node(4), Node(5)), Node(3, Node(6), Node(7)))
    solution.connect(root)
    assert root.next is None

    assert root.left and root.right
    assert root.left.next == root.right

    assert root.left.left and root.left.right
    assert root.right.left and root.right.right

    assert root.left.left.next == root.left.right
    assert root.left.right.next == root.right.left
    assert root.right.left.next == root.right.right
    assert root.right.right.next is None


def test_connect_empty():
    solution = Solution()
    root = None
    solution.connect(root)
    assert root is None
