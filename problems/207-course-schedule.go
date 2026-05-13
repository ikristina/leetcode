package problems

func canFinish(numCourses int, prerequisites [][]int) bool {

	// adjacent list graph representation
	graph := map[int][]int{}
	for _, c := range prerequisites {
		prereq, course := c[0], c[1]
		graph[prereq] = append(graph[prereq], course)
	}


	// use DFS to detect a cycle
	// 0 = unvisited
	// 1 = currently being visited (in the current DFS path)
	// 2 = fully processed (safe, no cycle found through here)

	state := []int{}
	for range numCourses {
		state = append(state, 0)
	}

	var dfs func(course int) bool

	dfs = func(course int) bool {
		if state[course] == 1 {
			return false
		}
		if state[course] == 2 {
			return true
		}
		state[course] = 1
		for _, neighbor := range graph[course] {
			if !dfs(neighbor) {
				return false
			}
		}
		state[course] = 2
		return true
	}

	for i := range numCourses {
		if !dfs(i) {
			return false
		}
	}

	return true
}
