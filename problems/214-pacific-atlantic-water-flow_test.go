package problems

import (
	"reflect"
	"testing"
)

func TestPacificAtlantic(t *testing.T) {
	tests := map[string]struct {
		heights [][]int
		want    [][]int
	}{
		"example 1": {
			heights: [][]int{
				{1, 2, 2, 3, 5},
				{3, 2, 3, 4, 4},
				{2, 4, 5, 3, 1},
				{6, 7, 1, 4, 5},
				{5, 1, 1, 2, 4},
			},
			want: [][]int{
				{0, 4},
				{1, 3},
				{1, 4},
				{2, 2},
				{3, 0},
				{3, 1},
				{4, 0},
			},
		},
		"example 2": {
			heights: [][]int{
				{1},
			},
			want: [][]int{
				{0, 0},
			},
		},
		"example 3": {
			heights: [][]int{
				{1, 2, 3},
				{8, 9, 4},
				{7, 6, 5},
			},
			want: [][]int{
				{0, 2},
				{1, 0},
				{1, 1},
				{1, 2},
				{2, 0},
				{2, 1},
				{2, 2},
			},
		},

	}
	for name, test := range tests {
		t.Run(name, func(t *testing.T) {
			got := pacificAtlantic(test.heights)
			if !reflect.DeepEqual(got, test.want) {
				t.Errorf("pacificAtlantic(%v) = %v, want %v", test.heights, got, test.want)
			}
		})
	}
}
