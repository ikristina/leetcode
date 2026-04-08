package problems

import "testing"

func TestLevelOrder(t *testing.T) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	expected := [][]int{{3}, {9, 20}, {15, 7}}
	result := levelOrder(root)
	if len(result) != len(expected) {
		t.Errorf("Expected %v, got %v", expected, result)
	}
	for i := range result {
		if len(result[i]) != len(expected[i]) {
			t.Errorf("Expected %v, got %v", expected[i], result[i])
		}
		for j := range result[i] {
			if result[i][j] != expected[i][j] {
				t.Errorf("Expected %v, got %v", expected[i][j], result[i][j])
			}
		}
	}
}

func BenchmarkLevelOrder(b *testing.B) {
	root := &TreeNode{Val: 3}
	root.Left = &TreeNode{Val: 9}
	root.Right = &TreeNode{Val: 20}
	root.Right.Left = &TreeNode{Val: 15}
	root.Right.Right = &TreeNode{Val: 7}
	for i := 0; i < b.N; i++ {
		levelOrder(root)
	}
}
