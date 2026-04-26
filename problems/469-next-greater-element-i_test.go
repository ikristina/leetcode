package problems

import (
	"slices"
	"testing"
)

func TestNextGreaterElementI(t *testing.T) {
	tests := map[string]struct {
		nums1 []int
		nums2 []int
		want  []int
	}{
		"case0": {
			nums1: []int{4, 1, 2},
			nums2: []int{1, 3, 4, 2},
			want:  []int{-1, 3, -1},
		},
		"case1": {
			nums1: []int{2, 4},
			nums2: []int{1, 2, 3, 4},
			want:  []int{3, -1},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := nextGreaterElement(tc.nums1, tc.nums2)
			if !slices.Equal(got, tc.want) {
				t.Errorf("wrong values, got: %v, want: %v", got, tc.want)
			}
		})
	}
}
