package problems

func cloneGraph(root *GraphNode) *GraphNode {
	if root == nil {
		return nil
	}

	clone := &GraphNode{Val: root.Val}
	visited := map[*GraphNode]*GraphNode{
		root: clone,
	}

	q := []*GraphNode{root}
	for len(q) > 0 {
		curr := q[0]
		q = q[1:]
		for _, n := range curr.Neighbors {
			if _, ok := visited[n]; !ok {
				visited[n] = &GraphNode{Val: n.Val}
				q = append(q, n)
			}
			visited[curr].Neighbors = append(visited[curr].Neighbors, visited[n])
		}
	}

	return clone
}
