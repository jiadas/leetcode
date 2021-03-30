package golang

func reverseList206(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	result := reverseList206(head.Next)
	head.Next = nil
	head.Next.Next = head
	return result
}
