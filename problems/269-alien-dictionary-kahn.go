package problems

// alienOrderKahn leetcode.com/problems/alien-dictionary
// hard
// #graph, #topological sort, #bfs, #kahn's algorithm

func alienOrderKahn(words []string) string {
	// inDegree tracks the number of incoming edges for each character.
	// adj is the adjacency list representing the graph.
	inDegree := map[byte]int{}
	adj := map[byte][]byte{}

	// Initialize inDegree for all unique characters present in the words.
	for _, w := range words {
		for i := 0; i < len(w); i++ {
			if _, ok := inDegree[w[i]]; !ok {
				inDegree[w[i]] = 0
			}
		}
	}

	// Build the graph by comparing adjacent words.
	for i := 0; i < len(words)-1; i++ {
		a, b := words[i], words[i+1]
		minLen := min(len(a), len(b))
		// If word A is a prefix of word B but is longer, it's an invalid order (e.g., "apple", "app").
		if len(a) > len(b) && a[:minLen] == b[:minLen] {
			return ""
		}
		// Find the first differing character to establish a directed edge.
		for j := 0; j < minLen; j++ {
			if a[j] != b[j] {
				adj[a[j]] = append(adj[a[j]], b[j])
				inDegree[b[j]]++
				break
			}
		}
	}

	// Kahn's Algorithm Step 1: Find all nodes with 0 in-degree.
	queue := []byte{}
	for c, deg := range inDegree {
		if deg == 0 {
			queue = append(queue, c)
		}
	}

	// Kahn's Algorithm Step 2: Process nodes in the queue.
	result := []byte{}
	for len(queue) > 0 {
		c := queue[0]
		queue = queue[1:]
		result = append(result, c)

		// For each neighbor, decrement its in-degree.
		for _, neighbor := range adj[c] {
			inDegree[neighbor]--
			// If in-degree becomes 0, add it to the queue.
			if inDegree[neighbor] == 0 {
				queue = append(queue, neighbor)
			}
		}
	}

	// If the result length doesn't match the number of unique characters, there's a cycle.
	if len(result) != len(inDegree) {
		return "" // cycle
	}
	return string(result)
}
