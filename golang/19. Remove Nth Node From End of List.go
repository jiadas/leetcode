package golang

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, f := head, head
	for i := 0; i < n; i++ {
		if f != nil {
			f = f.Next
		}
	}
	if f == nil {
		return p.Next
	}
	for f.Next != nil {
		p = p.Next
		f = f.Next
	}
	p.Next = p.Next.Next
	return head
}
