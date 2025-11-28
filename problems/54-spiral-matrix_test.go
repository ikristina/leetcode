package problems

import (
	"reflect"
	"testing"
)

func TestSpiralOrder(t *testing.T) {
	tests := map[string]struct {
		matrix [][]int
		want   []int
	}{
		"empty":       {matrix: [][]int{}, want: []int{}},
		"nil":         {matrix: nil, want: []int{}},
		"single row":  {matrix: [][]int{{1, 2, 3}}, want: []int{1, 2, 3}},
		"single col":  {matrix: [][]int{{1}, {2}, {3}, {4}}, want: []int{1, 2, 3, 4}},
		"1x1":         {matrix: [][]int{{42}}, want: []int{42}},
		"2x2":         {matrix: [][]int{{1, 2}, {3, 4}}, want: []int{1, 2, 4, 3}},
		"3x3":         {matrix: [][]int{{1, 2, 3}, {4, 5, 6}, {7, 8, 9}}, want: []int{1, 2, 3, 6, 9, 8, 7, 4, 5}},
		"rect wider":  {matrix: [][]int{{1, 2, 3, 4}, {5, 6, 7, 8}, {9, 10, 11, 12}}, want: []int{1, 2, 3, 4, 8, 12, 11, 10, 9, 5, 6, 7}},
		"rect taller": {matrix: [][]int{{1, 2}, {3, 4}, {5, 6}}, want: []int{1, 2, 4, 6, 5, 3}},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := spiralOrder(tt.matrix)
			if !reflect.DeepEqual(got, tt.want) {
				t.Fatalf("spiralOrder(%v) = %v, want %v", tt.matrix, got, tt.want)
			}
		})
	}
}
