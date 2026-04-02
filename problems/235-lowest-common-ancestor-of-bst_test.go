package problems

import "testing"

func TestLowestCommonAncestor(t *testing.T) {
	tests := map[string]struct {
		root *TreeNode
		p    *TreeNode
		q    *TreeNode
		want *TreeNode
	}{
		"test1": {
			root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}},
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 8},
			want: &TreeNode{Val: 6},
		},
		"test2": {
			root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}},
			p:    &TreeNode{Val: 2},
			q:    &TreeNode{Val: 4},
			want: &TreeNode{Val: 2},
		},
		"test3": {
			root: &TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}},
			p:    &TreeNode{Val: 3},
			q:    &TreeNode{Val: 5},
			want: &TreeNode{Val: 4},
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := lowestCommonAncestor(tt.root, tt.p, tt.q); got == nil || got.Val != tt.want.Val {
				t.Errorf("lowestCommonAncestor(%v, %v, %v) = %v, want %v", tt.root, tt.p, tt.q, got, tt.want)
			}
		})
	}
}

func BenchmarkLowestCommonAncestor(b *testing.B) {
	for b.Loop() {
		lowestCommonAncestor(&TreeNode{Val: 6, Left: &TreeNode{Val: 2, Left: &TreeNode{Val: 0}, Right: &TreeNode{Val: 4, Left: &TreeNode{Val: 3}, Right: &TreeNode{Val: 5}}}, Right: &TreeNode{Val: 8, Left: &TreeNode{Val: 7}, Right: &TreeNode{Val: 9}}}, &TreeNode{Val: 2}, &TreeNode{Val: 8})
	}
}
