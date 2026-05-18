import heapq
from typing import Optional, Sequence


class ListNode:
    def __init__(self, val=0, next=None):
        self.val = val
        self.next = next


class Solution:
    def mergeKLists(self, lists: Sequence[Optional[ListNode]]) -> Optional[ListNode]:

        h = []

        for i, list in enumerate(lists):
            if list:
                node = (list.val, i, list)
                heapq.heappush(h, node)

        dummy = ListNode(0)
        curr = dummy

        while h:
            val, i, node = heapq.heappop(h)
            curr.next = node
            curr = curr.next
            if node.next:
                heapq.heappush(h, (node.next.val, i, node.next))

        return dummy.next
