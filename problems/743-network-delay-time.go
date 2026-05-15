package problems

import "math"

func networkDelayTime(times [][]int, n, k int) int {

    // Step 1: Initialize distances to all nodes as infinity, except the source node (k).
    dist := map[int]int{}
    for i := 1; i <= n; i++ {
        dist[i] = math.MaxInt64
    }
    dist[k] = 0

    // Step 2: Build an adjacency list representation of the graph.
    graph := map[int][][]int{}
    for _, tt := range times {
        // tt[0] is source, tt[1] is target, tt[2] is weight (travel time).
        graph[tt[0]] = append(graph[tt[0]], []int{tt[1], tt[2]})
    }

    // Step 3: Use a queue to explore the graph (Dijkstra's Algorithm).
    // Note: A standard Dijkstra uses a Priority Queue for O(E log V) efficiency.
    // This implementation uses a simple queue (BFS style) but performs "relaxation".
    queue := []int{k}
    for len(queue) > 0 {
        node := queue[0]
        queue = queue[1:]

        for _, edge := range graph[node] {
            neighbor, travelTime := edge[0], edge[1]

            // "Relaxation" step: check if the path to 'neighbor' through 'node' is shorter
            // than the best path we've found so far.
            newTime := dist[node] + travelTime
            if newTime < dist[neighbor] {
                dist[neighbor] = newTime
                // If we found a better path, we need to re-explore this neighbor's neighbors.
                queue = append(queue, neighbor)
            }
        }
    }

	// Step 4: Find the maximum of all shortest paths.
	maxTime := 0
	for i := 1; i <= n; i++ {
	    time := dist[i]
	    // If any node is unreachable (still at infinity), we can't notify the whole network.
	    if time == math.MaxInt64 {
	        return -1
	    }
	    if time > maxTime {
	        maxTime = time
	    }
	}
	return maxTime
}
