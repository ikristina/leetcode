package problems

import "testing"

func TestIsValidBST(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		want bool
	}{
		"test1": {
			root: &TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}},
			want: true,
		},
		"test2": {
			root: &TreeNode{Val: 5, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 6}}},
			want: false,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := isValidBST(tt.root); got != tt.want {
				t.Errorf("isValidBST(%v) = %v, want %v", tt.root, got, tt.want)
			}
		})
	}
}

func BenchmarkIsValidBST(b *testing.B) {
	for b.Loop() {
		isValidBST(&TreeNode{Val: 2, Left: &TreeNode{Val: 1}, Right: &TreeNode{Val: 3}})
	}
}
