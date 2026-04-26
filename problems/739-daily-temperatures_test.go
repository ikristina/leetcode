package problems

import (
	"slices"
	"testing"
)

func TestDailyTemperatures(t *testing.T) {
	tests := map[string]struct {
		temperatures []int
		expected     []int
	}{
		"case 0": {
			temperatures: []int{73, 74, 75, 71, 69, 72, 76, 73},
			expected:     []int{1, 1, 4, 2, 1, 1, 0, 0},
		},
		"case 1": {
			temperatures: []int{30, 40, 50, 60},
			expected:     []int{1, 1, 1, 0},
		},
		"case 2": {
			temperatures: []int{30, 60, 90},
			expected:     []int{1, 1, 0},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := dailyTemperatures(tc.temperatures)
			if !slices.Equal(got, tc.expected) {
				t.Errorf("slices are not equal, got: %v, expected: %v", got, tc.expected)
			}
		})
	}
}
