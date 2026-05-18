package problems

import "testing"

func TestLastStoneWeight(t *testing.T) {
	tests := []struct {
		name     string
		input    []int
		expected int
	}{
		{
			name:     "example 1",
			input:    []int{2, 7, 4, 1, 8, 1},
			expected: 1,
		},
		{
			name:     "example 2",
			input:    []int{1},
			expected: 1,
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			result := lastStoneWeight(test.input)
			if result != test.expected {
				t.Errorf("expected %d, got %d", test.expected, result)
			}
		})
	}
}
