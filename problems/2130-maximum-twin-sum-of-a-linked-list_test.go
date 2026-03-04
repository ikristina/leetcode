package problems

import "testing"

func TestTwinSum(t *testing.T) {
	tests := map[string]struct {
		head *ListNode
		want int
	}{
		"test1": {
			head: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 5, Next: &ListNode{Val: 7}}}}, // 2->3->5->7
			want: 9,
		},

		"test2": {
			head: &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 7}}}}, // 4->5->6->7
			want: 11,
		},
	}

	for name, tt := range tests {
		t.Run(name, func(t *testing.T) {
			if got := TwinSum(tt.head); got != tt.want {
				t.Errorf("TwinSum(%v) = %v, want %v", tt.head, got, tt.want)
			}
		})
	}
}
