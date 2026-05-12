package problems

import "testing"

func TestBinarySearch(t *testing.T) {
	tests := map[string]struct {
		nums     []int
		target   int
		expected int
	}{
		"found": {
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   9,
			expected: 4,
		},
		"not found": {
			nums:     []int{-1, 0, 3, 5, 9, 12},
			target:   2,
			expected: -1,
		},

	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := binarySearch(test.nums, test.target)
			if result != test.expected {
				t.Errorf("Expected %d, got %d", test.expected, result)
			}
		})
	}
}
