package problems

import "testing"

func TestValidTree(t *testing.T) {
	tests := []struct {
		name     string
		n        int
		edges    [][]int
		expected bool
	}{
		{
			name:     "Example 1",
			n:        5,
			edges:    [][]int{{0, 1}, {0, 2}, {0, 3}, {1, 4}},
			expected: true,
		},
		{
			name:     "Example 2",
			n:        5,
			edges:    [][]int{{0, 1}, {1, 2}, {2, 3}, {1, 3}, {1, 4}},
			expected: false,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := validTree(test.n, test.edges)
			if result != test.expected {
				t.Errorf("Expected %v, got %v", test.expected, result)
			}
		})
	}
}
