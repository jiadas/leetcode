package golang

func isPalindrome234(head *ListNode) bool {
	if head == nil {
		return false
	}

	// 取一半链表可以用快慢指针的方法
	count := countLinkedList(head)
	rightHalf := head
	for i := 0; i < count/2; i++ {
		if rightHalf != nil {
			rightHalf = rightHalf.Next
		}
	}
	if count%2 == 1 && rightHalf != nil {
		rightHalf = rightHalf.Next
	}

	rightHalf = reverseListIteratively(rightHalf)
	leftHalf := head
	for leftHalf != nil && rightHalf != nil {
		if leftHalf.Val != rightHalf.Val {
			return false
		}
		leftHalf = leftHalf.Next
		rightHalf = rightHalf.Next
	}
	return true
}
