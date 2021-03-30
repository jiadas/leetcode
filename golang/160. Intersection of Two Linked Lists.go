package golang

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	countA := countLinkedList(headA)
	countB := countLinkedList(headB)
	short, long := headA, headB
	if countA > countB {
		short, long = long, short
	}

	for i := 0; i < absInt(countA-countB); i++ {
		if long != nil {
			long = long.Next
		}
	}

	for long != nil && short != nil {
		if long == short {
			return long
		}

		long = long.Next
		short = short.Next
	}
	return nil
}

func countLinkedList(head *ListNode) int {
	var count int
	walk := head
	for walk != nil {
		count++
		walk = walk.Next
	}
	return count
}
