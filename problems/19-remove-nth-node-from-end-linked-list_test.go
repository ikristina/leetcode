package problems

import (
	"reflect"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	tests := map[string]struct {
		head *ListNode
		n    int
		want *ListNode
	}{
		"test1": {
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}, // 1->2->3->4->5
			n:    2,
			want: &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5}}}}, // 1->2->3->5
		},
		"test2": {
			head: &ListNode{Val: 1}, // 1
			n:    1,
			want: nil, // empty list
		},
		"test3": {
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, // 1->2
			n:    1,
			want: &ListNode{Val: 1}, // 1
		},
		"test4": {
			head: &ListNode{Val: 1, Next: &ListNode{Val: 2}}, // 1->2
			n:    2,
			want: &ListNode{Val: 2}, // 2
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := RemoveNthFromEnd(tt.head, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveNthFromEnd(%v, %d) = %v, want %v", tt.head, tt.n, got, tt.want)
			}
		})
	}
}
