package problems

import "testing"

func TestKthSmallest(t *testing.T) {
	root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
	k := 2
	result := kthSmallest(root, k)
	if result != 2 {
		t.Errorf("Expected 2, got %d", result)
	}
}

func BenchmarkKthSmallest(b *testing.B) {
	for b.Loop() {
		root := &TreeNode{Val: 3, Left: &TreeNode{Val: 1, Right: &TreeNode{Val: 2}}, Right: &TreeNode{Val: 4}}
		k := 2
		kthSmallest(root, k)
	}
}
