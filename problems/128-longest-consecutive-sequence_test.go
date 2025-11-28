package problems

import (
	"testing"
)

func TestLongestConsecutive(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"case 0": {
			nums: []int{100, 4, 200, 1, 3, 2},
			want: 4,
		},
		"case 1": {
			nums: []int{0, 3, 7, 2, 5, 8, 4, 6, 0, 1},
			want: 9,
		},
		"case 2": {
			nums: []int{1, 0, 1, 2},
			want: 3,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := longestConsecutive(tc.nums)
			if got != tc.want {
				t.Errorf("wrong length. Want %d, got %d", tc.want, got)
			}
		})
	}
}
