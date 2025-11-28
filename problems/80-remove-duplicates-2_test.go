package problems

import "testing"

func Test_removeDuplicates(t *testing.T) {
	tests := map[string]struct {
		nums []int
		want int
	}{
		"test1": {
			nums: []int{1, 1, 1, 2, 2, 3},
			want: 5,
		},
		"test2": {
			nums: []int{0, 0, 1, 1, 1, 1, 2, 3, 3},
			want: 7,
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
			nums: []int{1, 1, 1, 1, 1},
			want: 2,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := removeDuplicates2(tt.nums); got != tt.want {
				t.Errorf("removeDuplicates() = %v, want %v", got, tt.want)
			}
		})
	}
}
