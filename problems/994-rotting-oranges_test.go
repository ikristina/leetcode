package problems

import "testing"

func TestOrangesRotting(t *testing.T) {
	tests := map[string]struct {
		grid    [][]int
		expected int
	}{
		"example1": {
			grid:    [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}},
			expected: 4,
		},
		"example2": {
			grid:    [][]int{{2, 1, 1}, {0, 1, 1}, {1, 0, 1}},
			expected: -1,
		},
		"example3": {
			grid:    [][]int{{0, 2}},
			expected: 0,
		},

	}
	for name, test := range tests {
		result := orangesRotting(test.grid)
		if result != test.expected {
			t.Errorf("Expected %d, got %d. Test: %s", test.expected, result, name)
		}
	}
}

func BenchmarkOrangesRotting(b *testing.B) {
	grid := [][]int{{2, 1, 1}, {1, 1, 0}, {0, 1, 1}}
	for b.Loop() {
		orangesRotting(grid)
	}
}
