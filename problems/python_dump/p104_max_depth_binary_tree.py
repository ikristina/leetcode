from collections import deque
from typing import Optional

from common_types import TreeNode


class Solution:
    def maxDepth(self, root: Optional[TreeNode]) -> int:
        if root is None:
            return 0

        queue = deque([root])
        depth = 0

        while queue:
            depth += 1
            level_size = len(queue)

            for _ in range(level_size):
                node = queue.popleft()
                if node.left:
                    queue.append(node.left)
                if node.right:
                    queue.append(node.right)

        return depth
