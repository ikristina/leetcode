import heapq
from typing import List


def leastInterval(tasks: List[str], n: int) -> int:
    tt = {}
    for t in tasks:
        if t not in tt:
            tt[t] = tasks.count(t)

    h = []
    for k, v in tt.items():
        heapq.heappush(h, (-v, k))

    cpuCount = 0

    while len(h):
        cycle = n + 1  # each round has n+1 slots
        used = 0
        temp = []

        for _ in range(cycle):
            if h:
                temp.append(heapq.heappop(h))
                used += 1

        for count, task in temp:
            if count + 1 < 0:  # remember counts are negative
                heapq.heappush(h, (count + 1, task))

        if h:
            cpuCount += cycle  # full round
        else:
            cpuCount += used  # last round, no idle padding

    return cpuCount
