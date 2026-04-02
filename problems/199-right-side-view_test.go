package problems

import (
	"reflect"
	"testing"
)

func TestRightSideView(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		want []int
	}{
		"test1": {
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}},
			want: []int{1, 3, 4},
		},
		"test2": {
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want: []int{1, 3},
		},
		"test3": {
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 3}},
			want: []int{1, 3, 4},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := rightSideView(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rightSideView(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}

func BenchmarkRightSideView(b *testing.B) {
	for b.Loop() {
		rightSideView(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3, Right: &TreeNode{Val: 4}}})
	}
}
