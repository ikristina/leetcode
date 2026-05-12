package problems

import "testing"

func TestNumberOfIslands(t *testing.T) {
	tests := map[string]struct {
		grid     [][]byte
		expected int
	}{
		"case 0": {
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
			expected: 1,
		},
		"case 1": {
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
			expected: 3,
		},
		"empty": {
			grid:     [][]byte{},
			expected: 0,
		},
	}
	for name, test := range tests {
		result := numberOfIslands(test.grid)
		if result != test.expected {
			t.Errorf("%s: Expected %d, got %d", name, test.expected, result)
		}
	}
}
