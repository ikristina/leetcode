package problems

import "testing"

func TestBuildTree(t *testing.T) {
    preorder := []int{3, 9, 20, 15, 7}
    inorder := []int{9, 3, 15, 20, 7}
    root := buildTree(preorder, inorder)
    if root == nil {
        t.Errorf("Expected non-nil root, got nil")
    }
    if root.Val != 3 {
        t.Errorf("Expected root value 3, got %d", root.Val)
    }
    if root.Left == nil {
        t.Errorf("Expected non-nil left child, got nil")
    }
    if root.Right == nil {
        t.Errorf("Expected non-nil right child, got nil")
    }
    if root.Left.Val != 9 {
        t.Errorf("Expected left child value 9, got %d", root.Left.Val)
    }
    if root.Right.Val != 20 {
        t.Errorf("Expected right child value 20, got %d", root.Right.Val)
    }
    if root.Right.Left.Val != 15 {
        t.Errorf("Expected left child value 15, got %d", root.Right.Left.Val)
    }
    if root.Right.Right.Val != 7 {
        t.Errorf("Expected right child value 7, got %d", root.Right.Right.Val)
    }
}

// BenchmarkBuildTree-10    	12492748	        95.71 ns/op	     160 B/op	       5 allocs/op
func BenchmarkBuildTree(b *testing.B) {
    preorder := []int{3, 9, 20, 15, 7}
    inorder := []int{9, 3, 15, 20, 7}
    for b.Loop() {
        buildTree(preorder, inorder)
    }
}
