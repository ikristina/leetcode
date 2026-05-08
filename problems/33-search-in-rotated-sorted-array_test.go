package problems

import "testing"

func TestSearch(t *testing.T) {
	testCases := []struct {
		nums   []int
		target int
		want   int
	}{
		{[]int{4, 5, 6, 7, 0, 1, 2}, 0, 4},
		{[]int{4, 5, 6, 7, 0, 1, 2}, 3, -1},
		{[]int{1}, 0, -1},
		{[]int{1}, 1, 0},
		{[]int{5, 1, 3}, 1, 1},
	}

	for _, tc := range testCases {
		got := search(tc.nums, tc.target)
		if got != tc.want {
			t.Errorf("search(%v, %d) = %d; want %d", tc.nums, tc.target, got, tc.want)
		}
	}
}
