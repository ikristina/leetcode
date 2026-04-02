package problems

import "testing"

func TestIsSubtree(t *testing.T) {
	tests := map[string]struct {
		tree    *TreeNode
		subtree *TreeNode
		want    bool
	}{
		"test 1": {
			tree:    &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}},
			subtree: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			want:    true,
		},
		"test 2": {
			tree:    &TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}}}, Right: &TreeNode{Val: 5}},
			subtree: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}},
			want:    false,
		},
		"test 3": {
			tree:    &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			subtree: &TreeNode{Val: 2},
			want:    true,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isSubtree(tt.tree, tt.subtree); got != tt.want {
				t.Errorf("isSubtree(%v, %v) = %v, want %v", tt.tree, tt.subtree, got, tt.want)
			}
		})
	}
}

func BenchmarkIsSubtree(b *testing.B) {
	for b.Loop() {
		isSubtree(&TreeNode{Val: 3, Left: &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 5}}, &TreeNode{Val: 4, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 2}})
	}
}
