package problems

import (
	"testing"
)

func TestAlienOrder(t *testing.T) {
	tests := map[string]struct {
		words   []string
		order   string
	}{
		"test1": {
			words:   []string{"xro","xma","per","prt","oxh","olv"},
			order:   "mrehlvopxat",
		},
		"test2": {
			words:   []string{"o","l","m","s"},
			order:   "smlo",
		},
		"test3": {
			words:   []string{"mdx","mars","avgd","dkae"},
			order:   "",
		},
		"test4": {
			words:   []string{"mdxok","mrolw","mroqs","kptz","klr","klon","zvef","zrsu","zzs","orm","oqt"},
			order:   "ozqrdxkwnfstvulemp",
		},
	}

	for name, test := range tests {
		got := alienOrder(test.words)
		graph := buildGraph(test.words)
		if test.order == "" {
			if got != "" {
				t.Errorf("%s: got %s, want %s", name, got, test.order)
			}
			continue
		}
		if !isValidOrder(graph, got) {
			t.Errorf("%s: got %s, want %s", name, got, test.order)
		}
	}
}

func isValidOrder(graph map[byte][]byte, result string) bool {
    if result == "" {
        return false
    }
    pos := map[byte]int{}
    for i := 0; i < len(result); i++ {
        pos[result[i]] = i
    }
    for src, dsts := range graph {
        for _, dst := range dsts {
            if pos[src] >= pos[dst] {
                return false
            }
        }
    }
    return true
}
