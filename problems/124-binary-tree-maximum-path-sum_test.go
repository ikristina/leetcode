package problems

import "testing"

func TestMaxPathSum(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		want int
	}{
		"simple": {
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}},
			want: 6,
		},
		"negative": {
			root: &TreeNode{Val: -1, Left: &TreeNode{Val: -2}, Right: &TreeNode{Val: -3}},
			want: -1,
		},
		"complex": {
			root: &TreeNode{Val: 1, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 4}}, Right: &TreeNode{Val: 5}},
			want: 12,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			if got := maxPathSum(tc.root); got != tc.want {
				t.Errorf("Expected %d, got %d", tc.want, got)
			}
		})
	}
}

func BenchmarkMaxPathSum(b *testing.B) {
	root := &TreeNode{Val: 1, Left: &TreeNode{Val: 2}, Right: &TreeNode{Val: 3}}
	for b.Loop() {
		maxPathSum(root)
	}
}
