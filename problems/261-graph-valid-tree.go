package problems

func validTree(n int, edges [][]int) bool {
	if len(edges) != n - 1 {
		return false
	}
	// create adjacency list representation of graph

	graph := make(map[int][]int)
	for i := range n {
		graph[i] = []int{}
	}
	for _, edge := range edges {
		graph[edge[0]] = append(graph[edge[0]], edge[1])
		graph[edge[1]] = append(graph[edge[1]], edge[0])
	}
	// dfs
	stack := []int{0}
	visited := make(map[int]bool)
	visited[0] = true

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		for _, neighbor := range graph[node] {
			if visited[neighbor] == false {
				stack = append(stack, neighbor)
				visited[neighbor] = true
			}
		}
	}

	return len(visited) == n
}
