from typing import Optional

from common_types import Node


class Solution:
    def connect(self, root: "Optional[Node]") -> "Optional[Node]":

        if root is None:
            return None
        # from collections import deque

        # q = deque([root])

        # while len(q):

        #     qlen = len(q)

        #     for i in range(qlen):
        #         node = q.popleft()
        #         if i < qlen - 1:
        #             node.next = q[0]

        #         if node.left:
        #             q.append(node.left)
        #         if node.right:
        #             q.append(node.right)

        # linked list approach:
        leftmost = root

        while leftmost and leftmost.left:
            head = leftmost
            while head and head.left and head.right:
                head.left.next = head.right
                if head.next:
                    head.right.next = head.next.left
                head = head.next
            leftmost = leftmost.left

        return root
