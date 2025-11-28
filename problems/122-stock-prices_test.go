package problems

import (
	"testing"
)

func Test122(t *testing.T) {
	tests := map[string]struct {
		input []int
		want  int
	}{
		"Example 1": {
			input: []int{7, 1, 5, 3, 6, 4},
			want:  7,
		},
		"Example 2": {
			input: []int{1, 2, 3, 4, 5},
			want:  4,
		},
		"Example 3": {
			input: []int{7, 6, 4, 3, 1},
			want:  0,
		},
		"Example 4": {
			input: []int{6, 1, 3, 2, 4, 7},
			want:  7,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := maxProfit2(tc.input)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
