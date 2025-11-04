package problems

import "testing"

func TestCopyRandomList(t *testing.T) {
	tests := map[string]struct {
		input    *Node
		expected *Node
	}{
		"empty node": {
			input:    nil,
			expected: nil,
		},
		"single node": {
			input:    &Node{Val: 1},
			expected: &Node{Val: 1},
		},
		"two nodes with random": {
			input: &Node{
				Val:  1,
				Next: &Node{Val: 2},
			},
			expected: &Node{
				Val:  1,
				Next: &Node{Val: 2},
			},
		},
		"three nodes with random pointers": {
			input: func() *Node {
				n1 := &Node{Val: 1}
				n2 := &Node{Val: 2}
				n3 := &Node{Val: 3}
				n1.Next = n2
				n2.Next = n3
				n1.Random = n3
				n2.Random = n1
				n3.Random = n2
				return n1
			}(),
			expected: func() *Node {
				n1 := &Node{Val: 1}
				n2 := &Node{Val: 2}
				n3 := &Node{Val: 3}
				n1.Next = n2
				n2.Next = n3
				n1.Random = n3
				n2.Random = n1
				n3.Random = n2
				return n1
			}(),
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			got := copyRandomList(tt.input)
			if !compareNodes(got, tt.expected) {
				t.Errorf("CopyRandom() = %v, want %v", got, tt.expected)
			}
		})
	}
}

// Helper function to compare two linked lists with random pointers
func compareNodes(a, b *Node) bool {
	if a == nil && b == nil {
		return true
	}
	if a == nil || b == nil {
		return false
	}

	nodeMap := make(map[*Node]*Node)
	curr1, curr2 := a, b

	// First pass: check values and build node mapping
	for curr1 != nil && curr2 != nil {
		if curr1.Val != curr2.Val {
			return false
		}
		nodeMap[curr1] = curr2
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	if curr1 != nil || curr2 != nil {
		return false
	}

	// Second pass: check random pointers
	curr1, curr2 = a, b
	for curr1 != nil {
		if curr1.Random == nil && curr2.Random != nil {
			return false
		}
		if curr1.Random != nil && curr2.Random == nil {
			return false
		}
		if curr1.Random != nil && nodeMap[curr1.Random] != curr2.Random {
			return false
		}
		curr1 = curr1.Next
		curr2 = curr2.Next
	}

	return true
}
