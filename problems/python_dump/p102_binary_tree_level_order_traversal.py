from typing import List, Optional

from common_types import TreeNode


class Solution:
    def levelOrder(self, root: Optional[TreeNode]) -> List[List[int]]:
        from collections import deque

        if not root:
            return []

        dq = deque([root])
        out = []

        while len(dq) > 0:
            level_len = len(dq)
            level = []

            for _ in range(level_len):
                node = dq.popleft()
                level.append(node.val)
                if node.left:
                    dq.append(node.left)
                if node.right:
                    dq.append(node.right)

            out.append(level)

        return out
