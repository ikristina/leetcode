package problems

import (
	"testing"
)

func TestKClosest(t *testing.T) {
	tests := map[string]struct {
		points [][]int
		k      int
		expected [][]int
	}{
		"simple": {
			points: [][]int{{1, 3}, {-2, 2}},
			k:      1,
			expected: [][]int{{-2, 2}},
		},
		"multiple": {
			points: [][]int{{3, 3}, {5, -1}, {-2, 4}},
			k:      2,
			expected: [][]int{{3, 3}, {-2, 4}},
		},
		"k_equals_len": {
			points: [][]int{{1, 1}, {2, 2}, {3, 3}},
			k:      3,
			expected: [][]int{{1, 1}, {2, 2}, {3, 3}},
		},
	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			result := kClosest(test.points, test.k)
			if !equal(result, test.expected) {
				t.Errorf("kClosest(%v, %d) = %v; expected %v", test.points, test.k, result, test.expected)
			}
		})
	}
}

func equal(a, b [][]int) bool {
	if len(a) != len(b) {
		return false
	}
	for _, el := range a {
		if !(contains(b, el)) {
			return false
		}
	}
	return true
}

func contains(a [][]int, el []int) bool {
	for _, e := range a {
		if e[0] == el[0] && e[1] == el[1] {
			return true
		}
	}
	return false
}
