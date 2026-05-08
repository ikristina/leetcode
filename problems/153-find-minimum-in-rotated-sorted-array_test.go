package problems

import "testing"

func TestFindMin(t *testing.T) {
    tests := []struct {
        nums []int
        want int
    }{
        {nums: []int{3, 4, 5, 1, 2}, want: 1},
        {nums: []int{4, 5, 6, 7, 0, 1, 2}, want: 0},
        {nums: []int{1}, want: 1},
        {nums: []int{2, 1}, want: 1},
    }

    for _, test := range tests {
        got := findMin(test.nums)
        if got != test.want {
            t.Errorf("findMin(%v) = %d, want %d", test.nums, got, test.want)
        }
    }
}
