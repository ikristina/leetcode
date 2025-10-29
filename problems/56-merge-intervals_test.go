package problems

import (
	"reflect"
	"testing"
)

func TestMergeIntervals(t *testing.T) {
	tests := map[string]struct {
		input    [][]int
		expected [][]int
	}{
		"case 1": {
			input:    [][]int{{1, 3}, {2, 6}, {8, 10}, {15, 18}},
			expected: [][]int{{1, 6}, {8, 10}, {15, 18}},
		},
		"case 2": {
			input:    [][]int{{1, 4}, {4, 5}},
			expected: [][]int{{1, 5}},
		},
		"case 3": {
			input:    [][]int{{1, 4}, {2, 3}},
			expected: [][]int{{1, 4}},
		},
		"case 4": {
			input:    [][]int{{1, 4}},
			expected: [][]int{{1, 4}},
		},
		"case 5": {
			input:    [][]int{},
			expected: [][]int{},
		},
	}
	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := mergeIntervals(tt.input)
			if !reflect.DeepEqual(got, tt.expected) {
				t.Errorf("%s: got %v, expected %v", name, got, tt.expected)
			}
		})
	}
}
