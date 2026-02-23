package problems

import "testing"

func TestFindDuplicate(t *testing.T) {
	tests := map[string]struct {
		n    []int
		want int
	}{
		"Base case 1": {n: []int{1, 3, 4, 2, 2}, want: 2},
		"Base case 2": {n: []int{3, 1, 3, 4, 2}, want: 3},
		"Example 4":   {n: []int{1, 1}, want: 1},
		"Example 5":   {n: []int{1, 1, 2}, want: 1},
		"Large n":     {n: []int{1, 1, 2, 3, 4}, want: 1},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := FindDuplicate(tt.n); got != tt.want {
				t.Errorf("FindDuplicate(%d) = %d, want %d", tt.n, got, tt.want)
			}
		})
	}
}
