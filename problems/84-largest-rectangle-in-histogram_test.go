package problems

import "testing"

func TestLargestRectangleArea(t *testing.T) {
	tests := map[string]struct{
		histogram []int
		expected  int
	}{
		"example1": {
			histogram: []int{2, 1, 5, 6, 2, 3},
			expected:  10,
		},
		"example2": {
			histogram: []int{2, 4},
			expected:  4,
		},
		"example3": {
			histogram: []int{},
			expected:  0,
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := largestRectangleArea(test.histogram)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
