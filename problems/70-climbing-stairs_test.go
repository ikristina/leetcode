package problems

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := []struct {
		name string
		n    int
		want int
	}{
		{"Base case 1", 1, 1},
		{"Base case 2", 2, 2},
		{"Example 3", 3, 3},
		{"Example 4", 4, 5},
		{"Example 5", 5, 8},
		{"Large n", 10, 89},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ClimbStairs(tt.n); got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
