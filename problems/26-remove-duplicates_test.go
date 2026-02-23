package problems

import "testing"

func TestRemoveDuplicates(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"example 1": {
			nums: []int{1, 1, 2},
			want: 2,
		},
		"example 2": {
			nums: []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4},
			want: 5,
		},
		"empty array": {
			nums: []int{},
			want: 0,
		},
		"single element": {
			nums: []int{1},
			want: 1,
		},
		"all same elements": {
			nums: []int{1, 1, 1, 1},
			want: 1,
		},
		"no duplicates": {
			nums: []int{1, 2, 3, 4, 5},
			want: 5,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := removeDuplicates(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
