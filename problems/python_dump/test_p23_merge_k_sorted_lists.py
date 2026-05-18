import p23_merge_k_sorted_lists as p23


def test_merge_k_sorted_lists():
    solution = p23.Solution()
    lists = [
        p23.ListNode(1, p23.ListNode(4, p23.ListNode(5))),
        p23.ListNode(1, p23.ListNode(3, p23.ListNode(4))),
        p23.ListNode(2, p23.ListNode(6)),
    ]
    result = solution.mergeKLists(lists)
    assert result.val == 1
    assert result.next.val == 1
    assert result.next.next.val == 2
    assert result.next.next.next.val == 3
    assert result.next.next.next.next.val == 4
    assert result.next.next.next.next.next.val == 4
    assert result.next.next.next.next.next.next.val == 5
    assert result.next.next.next.next.next.next.next.val == 6
    assert result.next.next.next.next.next.next.next.next is None
