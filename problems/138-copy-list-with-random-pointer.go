package problems

// Node...
// Definition for a Node.
type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

func copyRandomList(head *Node) *Node {
	// map of originalNode -> copyNode
	nm := map[*Node]*Node{nil: nil}

	// create copyNodes with values
	curr := head
	for curr != nil {
		copyCurr := &Node{Val: curr.Val}
		nm[curr] = copyCurr
		curr = curr.Next
	}

	// traverse the list again, find the copy, assign Next and Random
	curr = head
	for curr != nil {
		copyCurr := nm[curr]          // find the created copy
		copyCurr.Next = nm[curr.Next] // update next and random
		copyCurr.Random = nm[curr.Random]
		curr = curr.Next
	}

	return nm[head]
}
