package problems

import "testing"

func TestMaxAreaOfIsland(t *testing.T) {
	tests := map[string]struct {
		grid      [][]int
		expected  int
	}{
		"example": {
			grid: [][]int{
				{1, 1, 0, 0, 0},
				{1, 1, 0, 0, 0},
				{0, 0, 0, 1, 1},
				{0, 0, 0, 1, 1},
			},
			expected: 4,
		},
		"": {
			grid: [][]int{
			},
			expected: 0,
		},
		"single": {
			grid: [][]int{
				{1},
			},
			expected: 1,
		},
	}

	for name, test := range tests {
		grid := test.grid
		expected := test.expected
		result := maxAreaOfIsland(grid)
		if result != expected {
			t.Errorf("%s: Expected %d, got %d", name, expected, result)
		}
	}
}
