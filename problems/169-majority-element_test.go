package problems

import (
	"testing"
)

func Test_majorityElement(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"Example 1": {
			nums: []int{3, 2, 3},
			want: 3,
		},
		"Example 2": {
			nums: []int{2, 2, 1, 1, 1, 2, 2},
			want: 2,
		},
		"Single element": {
			nums: []int{1},
			want: 1,
		},
		"All same elements": {
			nums: []int{5, 5, 5, 5},
			want: 5,
		},
		"Negative numbers": {
			nums: []int{-1, -1, 2, -1},
			want: -1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := majorityElement(tt.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
