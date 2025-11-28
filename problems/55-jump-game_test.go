package problems

import "testing"

func TestCanJump(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want bool
	}{
		"example 1": {
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		"example 2": {
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		"single element": {
			nums: []int{0},
			want: true,
		},
		"empty array": {
			nums: []int{},
			want: true,
		},
		"can't jump from start": {
			nums: []int{0, 1, 2},
			want: false,
		},
		"large jump at start": {
			nums: []int{5, 0, 0, 0, 0, 1},
			want: true,
		},
		"stuck in middle": {
			nums: []int{1, 0, 1, 0},
			want: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}
