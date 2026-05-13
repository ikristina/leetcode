from typing import List


def canFinish(numCourses: int, prerequisites: List[List[int]]) -> bool:
    graph = {i: [] for i in range(numCourses)}  # adjacency list graph representation
    for course, prereq in prerequisites:
        graph[prereq].append(course)

    # use DFS to detect a cycle
    # 0 = unvisited
    # 1 = currently being visited (in the current DFS path)
    # 2 = fully processed (safe, no cycle found through here)

    state = [0] * numCourses

    def dfs(course):
        if state[course] == 1:
            return False
        if state[course] == 2:
            return True

        state[course] = 1
        for neighbor in graph[course]:
            if not dfs(neighbor):
                return False

        state[course] = 2
        return True

    return all(dfs(i) for i in range(numCourses))
