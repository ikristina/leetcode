package problems

import (
	"reflect"
	"testing"
)

func TestSortColors(t *testing.T) {
	tests := map[string]struct {
		input    []int
		expected []int
	}{
		"test1": {
			input:    []int{2, 0, 2, 1, 1, 0},
			expected: []int{0, 0, 1, 1, 2, 2},
		},
		"test2": {
			input:    []int{2, 0, 1},
			expected: []int{0, 1, 2},
		},
		"test3": {
			input:    []int{0, 1, 2},
			expected: []int{0, 1, 2},
		},
		"test4": {
			input:    []int{2, 1, 0},
			expected: []int{0, 1, 2},
		},
		"test5": {
			input:    []int{1, 2, 0},
			expected: []int{0, 1, 2},
		},
		"test6": {
			input:    []int{2, 1},
			expected: []int{1, 2},
		},
		"test7": {
			input:    []int{1, 2},
			expected: []int{1, 2},
		},
		"test8": {
			input:    []int{1, 0},
			expected: []int{0, 1},
		},
		"test9": {
			input:    []int{0, 0, 0},
			expected: []int{0, 0, 0},
		},
		"test10": {
			input:    []int{1, 1, 1},
			expected: []int{1, 1, 1},
		},
		"test11": {
			input:    []int{2, 2, 2},
			expected: []int{2, 2, 2},
		},
	}

	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			SortColors(test.input)
			if !reflect.DeepEqual(test.input, test.expected) {
				t.Errorf("SortColors(%v) = %v, want %v", test.input, test.input, test.expected)
			}
		})
	}
}

func BenchmarkSortColors(b *testing.B) {
	for b.Loop() {
		SortColors([]int{2, 0, 2, 1, 1, 0})
	}
}
