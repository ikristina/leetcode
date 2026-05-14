package problems

import "fmt"

func alienOrder(words []string) string {

	result := ""

	graph := buildGraph(words)
	if graph == nil {
		return result
	}

    visited := map[byte]int{} // 0, 1, 2

    var dfs func(node byte) bool

    dfs = func(node byte) bool {
        if visited[node] == 1 {
            return false
        }
        if visited[node] == 2 {
            return true
        }
        visited[node] = 1
        for _, n := range graph[node] {
            if !dfs(n) {
                return false
            }
        }
        visited[node] = 2
        result = result + string(rune(node))
        return true
    }

    chars := map[byte]struct{}{}
    for _, w := range words {
        for _, s := range w {

            chars[byte(s)] = struct{}{}
        }
    }
    fmt.Println(graph)
    for c := range chars {
        if !dfs(c) {
            return ""
        }
    }


	return result

}

func buildGraph(words []string) map[byte][]byte {
	graph := map[byte][]byte{}
	for i := 0; i < len(words)-1; i++ {
		a := words[i]
		b := words[i+1]

		for j := 0; j < len(a); j++{
			if j < len(b) {
			    if a[j] == b[j] {
			        continue
			    }
				graph[a[j]] = append(graph[a[j]], b[j])
				break
			}
			if j == len(b) && len(a) > len(b) {
			    return nil
			}
		}
	}
	return graph
}
