package golang

func removeElements(head *ListNode, val int) *ListNode {
	if head == nil {
		return nil
	}

	result := head
	for result != nil && result.Val == val {
		result = result.Next
	}

	if result == nil {
		return nil
	}

	pre, cur := result, result.Next
	for cur != nil {
		if cur.Val == val {
			pre.Next = cur.Next
		} else {
			pre = cur
		}
		cur = cur.Next
	}
	return result
}
