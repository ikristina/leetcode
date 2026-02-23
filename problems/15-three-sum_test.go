package problems

import (
	"reflect"
	"testing"
)

func TestThreeSum(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected [][]int
	}{
		"test1": {
			input:    []int{-1, 0, 1, 2, -1, -4},
			expected: [][]int{{-1, -1, 2}, {-1, 0, 1}},
		},
		"test2": {
			input:    []int{},
			expected: [][]int{},
		},
		"test3": {
			input:    []int{0},
			expected: [][]int{},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := ThreeSum(test.input)
			if !reflect.DeepEqual(result, test.expected) {
				t.Errorf("ThreeSum(%v) = %v, want %v", test.input, result, test.expected)
			}
		})
	}
}

func BenchmarkThreeSum(b *testing.B) {
	for b.Loop() {
		ThreeSum([]int{-1, 0, 1, 2, -1, -4})
	}
}
