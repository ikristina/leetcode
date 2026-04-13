package problems

import "testing"

func TestWaysToMakeFair(t *testing.T) {
	tests := map[string]struct {
		arr  []int
		want int
	}{
		"example1": {[]int{2, 1, 6, 4}, 1},
		"example2": {[]int{1, 1, 1}, 3},
		"example3": {[]int{1, 2, 3}, 0},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := waysToMakeFair(tc.arr); got != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, got)
			}
		})
	}
}

func BenchmarkWaysToMakeFair(b *testing.B) {
	arr := []int{}
	for b.Loop() {
		waysToMakeFair(arr)
	}
}
