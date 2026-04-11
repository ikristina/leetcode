package problems

import (
	"reflect"
	"testing"
)

func TestAverageOfLevels(t *testing.T) {
	tests := map[string]struct {
		root   *TreeNode
		want   []float64
	}{
		"single node": {
			root:   &TreeNode{Val: 1},
			want:   []float64{1},
		},
		"two levels": {
			root:   &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want:   []float64{3, 14.5, 11},
		},
		"three levels": {
			root:   &TreeNode{Val: 5, Left: &TreeNode{Val: 3, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 8}}},
			want:   []float64{5, 5, 4.75},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := averageOfLevels(tc.root)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("averageOfLevels(%v) = %v, want %v", tc.root, got, tc.want)
			}
		})
	}
}
