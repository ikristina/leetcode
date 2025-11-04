package problems

import "testing"

func TestCanJump(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{
			name: "example 1",
			nums: []int{2, 3, 1, 1, 4},
			want: true,
		},
		{
			name: "example 2",
			nums: []int{3, 2, 1, 0, 4},
			want: false,
		},
		{
			name: "single element",
			nums: []int{0},
			want: true,
		},
		{
			name: "empty array",
			nums: []int{},
			want: true,
		},
		{
			name: "can't jump from start",
			nums: []int{0, 1, 2},
			want: false,
		},
		{
			name: "large jump at start",
			nums: []int{5, 0, 0, 0, 0, 1},
			want: true,
		},
		{
			name: "stuck in middle",
			nums: []int{1, 0, 1, 0},
			want: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.nums); got != tt.want {
				t.Errorf("canJump() = %v, want %v", got, tt.want)
			}
		})
	}
}