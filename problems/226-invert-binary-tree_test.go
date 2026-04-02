package problems

import (
	"reflect"
	"testing"
)

func TestInvertBinaryTree(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		want *TreeNode
	}{
		"test1": {
			root: &TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}},
			want: &TreeNode{Val: 4, Left: &TreeNode{Val: 7, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}}},
		},
		"test2": {
			root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			want: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 1}},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := mirrorBinaryTree(tt.root); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mirrorBinaryTree(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}

func BenchmarkInvertBinaryTree(b *testing.B) {
	for b.Loop() {
		mirrorBinaryTree(&TreeNode{Val: 4, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}}, Right: &TreeNode{Val: 7, Left: &TreeNode{Val: 6}, Right: &TreeNode{Val: 9}}})
	}
}
