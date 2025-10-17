package problems

import "testing"

func TestMaxProfit(t *testing.T) {
	tests := []struct {
		name     string
		prices   []int
		expected int
	}{
		{
			name:     "Example 1",
			prices:   []int{7, 1, 5, 3, 6, 4},
			expected: 5,
		},
		{
			name:     "Example 2",
			prices:   []int{7, 6, 4, 3, 1},
			expected: 0,
		},
		{
			name:     "Empty array",
			prices:   []int{},
			expected: 0,
		},
		{
			name:     "Single element",
			prices:   []int{1},
			expected: 0,
		},
		{
			name:     "Two elements increasing",
			prices:   []int{1, 2},
			expected: 1,
		},
		{
			name:     "Two elements decreasing",
			prices:   []int{2, 1},
			expected: 0,
		},
		{
			name:     "Multiple peaks and valleys",
			prices:   []int{3, 2, 6, 1, 2, 5, 4},
			expected: 4,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := maxProfit(tt.prices)
			if got != tt.expected {
				t.Errorf("maxProfit() = %v, want %v", got, tt.expected)
			}
		})
	}
}
