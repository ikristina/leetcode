package problems

import (
	"slices"
	"testing"
)

func TestMerge(t *testing.T) {
	got := Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	want := []int{1, 2, 2, 3, 5, 6}
	if !slices.Equal(got, want) {
		t.Fatalf("Got: %v, Want: %v", got, want)
	}
}

func BenchmarkMerge(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = Merge([]int{1, 2, 3, 0, 0, 0}, 3, []int{2, 5, 6}, 3)
	}
}
