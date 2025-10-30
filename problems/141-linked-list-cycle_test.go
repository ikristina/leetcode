package problems

import "testing"

func TestHasCycle(t *testing.T) {
	tests := map[string]struct {
		head     *ListNode
		expected bool
	}{
		"no cycle": {
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: nil}},
			expected: false,
		},
		"cycle exists": {
			head: func() *ListNode {
				node1 := &ListNode{Val: 1}
				node2 := &ListNode{Val: 2}
				node3 := &ListNode{Val: 3}
				node1.Next = node2
				node2.Next = node3
				node3.Next = node1 // Creates a cycle
				return node1
			}(),
			expected: true,
		},
		"single node no cycle": {
			head:     &ListNode{Val: 1, Next: nil},
			expected: false,
		},
		"single node with cycle": {
			head:     &ListNode{Val: 1, Next: &ListNode{Val: 1, Next: nil}},
			expected: false,
		},
		"empty list": {
			head:     nil,
			expected: false,
		},
	}

	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			result := hasCycle(tc.head)
			if result != tc.expected {
				t.Errorf("expected %v, got %v", tc.expected, result)
			}
		})
	}
}
