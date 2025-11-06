package problems

import "testing"

func TestFindKthLargest(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		want int
	}{
		"case 1": {
			nums: []int{3, 2, 1, 5, 6, 4},
			k:    2,
			want: 5,
		},
		"case 2": {
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6},
			k:    4,
			want: 4,
		},
		"case 3": {
			nums: []int{1},
			k:    1,
			want: 1,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := findKthLargest(tc.nums, tc.k)
			if got != tc.want {
				t.Errorf("got: %v, want: %v", got, tc.want)
			}
		})
	}
}
