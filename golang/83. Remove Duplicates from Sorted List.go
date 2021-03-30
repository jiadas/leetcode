package golang

func deleteDuplicates(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		if fast.Next.Val == fast.Val {
			fast = fast.Next
			continue
		}

		fast = fast.Next
		slow.Next.Val = fast.Val
		slow = slow.Next
	}
	if slow != nil {
		slow.Next = nil
	}
	return head
}
