package problems

import "testing"

func TestMaxDepth(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		want int
	}{
		"test1": {
			root: &TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}},
			want: 3,
		},
		"test2": {
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want: 2,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := maxDepth(tt.root); got != tt.want {
				t.Errorf("maxDepth(%v) = %d, want %d", tt.root, got, tt.want)
			}
		})
	}
}

func BenchmarkMaxDepth(b *testing.B) {
	for b.Loop() {
		maxDepth(&TreeNode{Val: 3, Left: &TreeNode{Val: 9}, Right: &TreeNode{Val: 20, Left: &TreeNode{Val: 15}, Right: &TreeNode{Val: 7}}})
	}
}
