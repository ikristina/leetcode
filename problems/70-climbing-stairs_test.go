package problems

import "testing"

func TestClimbStairs(t *testing.T) {
	tests := map[string]struct {
		n    int
		want int
	}{
		"Base case 1": {n: 1, want: 1},
		"Base case 2": {n: 2, want: 2},
		"Example 3":   {n: 3, want: 3},
		"Example 4":   {n: 4, want: 5},
		"Example 5":   {n: 5, want: 8},
		"Large n":     {n: 10, want: 89},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := ClimbStairs(tt.n); got != tt.want {
				t.Errorf("ClimbStairs(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
