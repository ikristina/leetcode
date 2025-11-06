package problems

// ListNode definition would be here if not already defined

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	prev := &ListNode{}
	curr := prev

	carryOver := 0

	for l1 != nil || l2 != nil || carryOver > 0 {

		sum := carryOver

		if l1 != nil {
			sum += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			sum += l2.Val
			l2 = l2.Next
		}
		curr.Next = &ListNode{Val: sum % 10}
		curr = curr.Next
		carryOver = sum / 10

	}

	return prev.Next
}
