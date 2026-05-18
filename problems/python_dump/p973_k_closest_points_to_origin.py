import heapq
from typing import List


class Solution:
    def kClosest(self, points: List[List[int]], k: int) -> List[List[int]]:
        heap = []
        for point in points:
            dist = point[0] ** 2 + point[1] ** 2
            heapq.heappush(heap, (dist, point))
        result = []
        for _ in range(k):
            result.append(heapq.heappop(heap)[1])
        return result
