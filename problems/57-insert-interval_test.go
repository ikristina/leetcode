package problems

import (
	"reflect"
	"testing"
)

func TestInsertInterval(t *testing.T) {
	tests := []struct {
		name     string
		existing [][]int
		new      []int
		want     [][]int
	}{
		{
			name:     "empty intervals",
			existing: [][]int{},
			new:      []int{1, 3},
			want:     [][]int{{1, 3}},
		},
		{
			name:     "insert at beginning",
			existing: [][]int{{4, 6}, {8, 10}},
			new:      []int{1, 2},
			want:     [][]int{{1, 2}, {4, 6}, {8, 10}},
		},
		{
			name:     "insert at end",
			existing: [][]int{{1, 3}, {6, 9}},
			new:      []int{10, 12},
			want:     [][]int{{1, 3}, {6, 9}, {10, 12}},
		},
		{
			name:     "merge intervals",
			existing: [][]int{{1, 3}, {6, 9}},
			new:      []int{2, 7},
			want:     [][]int{{1, 9}},
		},
		{
			name:     "merge multiple intervals",
			existing: [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
			new:      []int{4, 9},
			want:     [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name:     "no overlap",
			existing: [][]int{{1, 2}, {4, 5}, {7, 8}},
			new:      []int{3, 3},
			want:     [][]int{{1, 2}, {3, 3}, {4, 5}, {7, 8}},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := insertInterval(tt.existing, tt.new)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("insertInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
