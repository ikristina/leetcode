package problems

import (
	"slices"
	"testing"
)

func TestTwoSum(t *testing.T) {
	got := TwoSum([]int{2, 7, 11, 15}, 9)
	want := []int{0, 1}
	if !slices.Equal(got, want) {
		t.Fatalf("Got: %v, Want: %v", got, want)
	}
}

func BenchmarkTwoSum(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = TwoSum([]int{2, 7, 11, 15}, 9)
	}
}
