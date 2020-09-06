package golang

func detectCycle(head *ListNode) *ListNode {
	fast, slow := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
		if fast == slow {
			slow = head
			for fast != slow {
				fast = fast.Next
				slow = slow.Next
			}
			return slow
		}
	}
	return nil
}
