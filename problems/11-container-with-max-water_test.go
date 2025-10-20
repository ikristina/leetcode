package problems

import "testing"

func TestMaxArea(t *testing.T) {
	tests := []struct {
		name     string
		height   []int
		expected int
	}{
		{
			name:     "Example 1",
			height:   []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			expected: 49,
		},
		{
			name:     "Example 2",
			height:   []int{1, 1},
			expected: 1,
		},
		{
			name:     "Example 3",
			height:   []int{0, 2},
			expected: 0,
		},
		{
			name:     "Example 4",
			height:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "Empty array",
			height:   []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			height:   []int{1},
			expected: 0,
		},
		{
			name:     "All same height",
			height:   []int{4, 4, 4, 4},
			expected: 12,
		},
		{
			name:     "Increasing heights",
			height:   []int{1, 2, 3, 4, 5},
			expected: 6,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxArea(tt.height)
			if got != tt.expected {
				t.Errorf("maxArea() = %v, want %v", got, tt.expected)
			}
		})
	}
}
