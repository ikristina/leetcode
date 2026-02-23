package problems

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"Example 1": {
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		"Example 2": {
			nums: []int{-2, 0, -1},
			want: 0,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
