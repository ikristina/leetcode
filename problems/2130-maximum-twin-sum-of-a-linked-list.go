package problems

// TwinSum Maximum Twin Sum of a Linked List
// Time Complexity: O(n)
// Space Complexity: O(1)
// #linked list, #fast and slow pointers
func TwinSum(head *ListNode) int {
	dummy := &ListNode{Next: head}

	current1 := head
	current2 := head

	// find middle
	for current1 != nil && current1.Next != nil {
		current2 = current2.Next
		current1 = current1.Next.Next
	}

	// reverse second half of the linked list
	var prev *ListNode = nil
	current := current2

	for current != nil {
		next := current.Next // store next node
		current.Next = prev  // reverse pointer
		prev = current       // move prev forward
		current = next
	}

	// reversed second half. prev is the new head
	s := 0
	current1 = dummy.Next
	current2 = prev
	for current1 != nil && current2 != nil {
		if s < current1.Val+current2.Val {
			s = current1.Val + current2.Val
		}
		current1 = current1.Next
		current2 = current2.Next
	}

	return s
}
