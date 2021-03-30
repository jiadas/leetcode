package cci

func kthToLast(head *ListNode, k int) int {
	if head == nil {
		return 0
	}

	back, front := head, head
	for i := 0; i < k; i++ {
		if front == nil {
			return 0
		}
		front = front.Next
	}
	for front != nil {
		back = back.Next
		front = front.Next
	}
	return back.Val
}
