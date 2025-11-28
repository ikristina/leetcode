package problems

import "testing"

func TestMaxProduct(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{
			name: "Example 1",
			nums: []int{2, 3, -2, 4},
			want: 6,
		},
		{
			name: "Example 2",
			nums: []int{-2, 0, -1},
			want: 0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxProduct(tt.nums); got != tt.want {
				t.Errorf("maxProduct() = %v, want %v", got, tt.want)
			}
		})
	}
}
