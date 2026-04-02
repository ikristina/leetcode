package problems

import "testing"

func TestDiameterOfBinaryTree(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		want int
	}{
		"test1": {
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}},
			want: 3,
		},
		"test2": {
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}},
			want: 1,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := diameterOfBinaryTree(tt.root); got != tt.want {
				t.Errorf("diameterOfBinaryTree(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}

func BenchmarkDiameterOfBinaryTree(b *testing.B) {
	for b.Loop() {
		diameterOfBinaryTree(&TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 4}, Right: &TreeNode{Val: 5}}, Right: &TreeNode{Val: 3}})
	}
}
