package problems

import "testing"

func TestRotate(t *testing.T) {
	tests := map[string]struct {
		nums []int
		k    int
		want []int
	}{
		"rotate by 3": {[]int{1, 2, 3, 4, 5, 6, 7}, 3, []int{5, 6, 7, 1, 2, 3, 4}},
		"rotate by 2": {[]int{-1, -100, 3, 99}, 2, []int{3, 99, -1, -100}},
		"no rotation": {[]int{1, 2}, 0, []int{1, 2}},
		"full rotation": {[]int{1, 2, 3}, 3, []int{1, 2, 3}},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			nums := make([]int, len(tc.nums))
			copy(nums, tc.nums)
			rotate(nums, tc.k)
			for i, v := range tc.want {
				if nums[i] != v {
					t.Errorf("got %v, want %v", nums, tc.want)
					break
				}
			}
		})
	}
}