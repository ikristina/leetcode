package problems

import (
	"slices"
	"testing"
)

func TestTwoSum2(t *testing.T) {
	tests := map[string]struct {
		input  []int
		target int
		want   []int
	}{
		"[2,7,11,15]": {
			input:  []int{2, 7, 11, 15},
			target: 9,
			want:   []int{1, 2},
		},

		"[2,3,4]": {
			input:  []int{2, 3, 4},
			target: 6,
			want:   []int{1, 3},
		},

		"[-1, 0]": {
			input:  []int{-1, 0},
			target: -1,
			want:   []int{1, 2},
		},

		"[-5,-3,0,2,4,6,8]": {
			input:  []int{-5, -3, 0, 2, 4, 6, 8},
			target: 5,
			want:   []int{2, 7},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := twoSum2(tc.input, tc.target)
			if !slices.Equal(got, tc.want) {
				t.Errorf("failed test %s, got: %v, want: %v", name, got, tc.want)
			}
		})
	}
}
