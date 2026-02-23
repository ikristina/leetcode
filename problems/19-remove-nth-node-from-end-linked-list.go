package problems

// RemoveNthFromEnd solves LeetCode 19: Remove Nth Node From End of List
// Time Complexity: O(n)
// Space Complexity: O(1)
// #linked list, #two pointers
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {

	dummy := &ListNode{Next: head}

	p1, p2 := dummy, dummy

	for i := 0; i < n; i++ {
		p1 = p1.Next
	}

	for p1 != nil && p1.Next != nil {
		p1 = p1.Next
		p2 = p2.Next
	}

	p2.Next = p2.Next.Next

	return dummy.Next

}
