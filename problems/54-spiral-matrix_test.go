package problems

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := []struct {
		name   string
		matrix [][]int
		want   []int
	}{
		{"empty", [][]int{}, []int{}},
		{"nil", nil, []int{}},
		{"single row", [][]int{{1, 2, 3}}, []int{1, 2, 3}},
		{"single col", [][]int{{1}, {2}, {3}, {4}}, []int{1, 2, 3, 4}},
		{"1x1", [][]int{{42}}, []int{42}},
		{"2x2", [][]int{{1, 2}, {3, 4}}, []int{1, 2, 4, 3}},
		{"3x3", [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		{"rect wider", [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		{"rect taller", [][]int{{1, 2}, {3, 4}, {5, 6}}, []int{1, 2, 4, 6, 5, 3}},
	}

	for _, tt := range tests {
		tt := tt
		t.Run(tt.name, func(t *testing.T) {
			got := spiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("spiralOrder(%v) = %v, want %v", tt.matrix, got, tt.want)
			}
		})
	}
}
