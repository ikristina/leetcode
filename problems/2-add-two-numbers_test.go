package problems

import (
	"fmt"
	"testing"
)

func TestAddTwoNumbers(t *testing.T) {
	tests := map[string]struct {
		l1   *ListNode
		l2   *ListNode
		want *ListNode
	}{
		"case 1": {
			l1:   &ListNode{2, &ListNode{4, &ListNode{3, nil}}},
			l2:   &ListNode{5, &ListNode{6, &ListNode{4, nil}}},
			want: &ListNode{7, &ListNode{0, &ListNode{8, nil}}},
		},
		"case 2": {
			l1:   &ListNode{0, nil},
			l2:   &ListNode{0, nil},
			want: &ListNode{0, nil},
		},
		"case 3": {
			l1:   &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}}}}},
			l2:   &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{9, nil}}}},
			want: &ListNode{8, &ListNode{9, &ListNode{9, &ListNode{9, &ListNode{0, &ListNode{0, &ListNode{0, &ListNode{1, nil}}}}}}}},
		},
		"case 4": {
			l1:   &ListNode{2, &ListNode{4, &ListNode{9, nil}}},
			l2:   &ListNode{5, &ListNode{6, &ListNode{4, &ListNode{9, nil}}}},
			want: &ListNode{7, &ListNode{0, &ListNode{4, &ListNode{0, &ListNode{1, nil}}}}},
		},
	}
	for name, tc := range tests {
		t.Run(name, func(t *testing.T) {
			got := addTwoNumbers(tc.l1, tc.l2)
			if !compareLinkedLists(got, tc.want) {
				t.Errorf("got: %v, want: %v", linkedListToString(got), linkedListToString(tc.want))
			}
		})
	}
}

// Helper function to compare two linked lists
func compareLinkedLists(l1 *ListNode, l2 *ListNode) bool {
	for l1 != nil && l2 != nil {
		if l1.Val != l2.Val {
			return false
		}
		l1 = l1.Next
		l2 = l2.Next
	}
	return l1 == nil && l2 == nil
}

// Helper function to convert linked list to string for error messages
func linkedListToString(l *ListNode) string {
	if l == nil {
		return "nil"
	}
	result := "["
	for l != nil {
		result += fmt.Sprintf("%d", l.Val)
		if l.Next != nil {
			result += ","
		}
		l = l.Next
	}
	result += "]"
	return result
}
