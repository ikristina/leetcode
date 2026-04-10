package problems

import (
	"reflect"
	"testing"
)

func TestConnect(t *testing.T) {
	tests := map[string]struct {
		input  *TreeNode
		want   *TreeNode
	}{
		"single node": {
			input: &TreeNode{Val: 1},
			want:  &TreeNode{Val: 1},
		},
		"two nodes": {
			input: &TreeNode{Val: 1, Next: &TreeNode{Val: 2}},
			want:  &TreeNode{Val: 1, Next: &TreeNode{Val: 2}},
		},
		"three nodes": {
			input: &TreeNode{Val: 1, Next: &TreeNode{Val: 2, Next: &TreeNode{Val: 3}}},
			want:  &TreeNode{Val: 1, Next: &TreeNode{Val: 2, Next: &TreeNode{Val: 3}}},
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := connect(tc.input)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got %v, want %v", got, tc.want)
			}
		})
	}
}
