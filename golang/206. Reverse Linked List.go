package golang

func reverseListRecursively(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := reverseListRecursively(head.Next)
	head.Next.Next = head
	head.Next = nil
	return result
}

func reverseListIteratively(head *ListNode) *ListNode {
	var pre, cur *ListNode = nil, head
	for cur != nil {
		nextTemp := cur.Next
		cur.Next = pre
		pre = cur
		cur = nextTemp
	}
	return pre
}

// reverseListFirstTry 自己的思路，没有 reverseListIteratively 来的简洁
func reverseListFirstTry(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	pre, cur, curNext := head, head.Next, head.Next.Next
	for cur != nil {
		cur.Next = pre
		pre = cur
		cur = curNext
		if cur != nil {
			curNext = cur.Next
		}
	}
	head.Next = nil
	return pre
}
