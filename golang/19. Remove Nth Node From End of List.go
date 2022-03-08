package golang

// removeNthFromEnd 参数提示：
// 链表中结点的数目为 sz
// 1 <= sz <= 30
// 0 <= Node.val <= 100
// 1 <= n <= sz
// https://labuladong.gitee.io/algo/2/22/59/
func removeNthFromEnd(head *ListNode, n int) *ListNode {
	p, f := head, head
	// 快指针先前进 n 步
	for i := 0; i < n; i++ {
		if f != nil {
			f = f.Next
		}
	}
	if f == nil {
		// 如果此时快指针走到头了，
		// 说明倒数第 n 个节点就是第一个节点
		return p.Next
	}
	// 让慢指针和快指针同步向前
	for f.Next != nil {
		p = p.Next
		f = f.Next
	}
	// slow.next 就是倒数第 n 个节点，删除它
	p.Next = p.Next.Next
	return head
}

// https://labuladong.gitee.io/algo/2/18/17/160
func findFromEnd(head *ListNode, n int) *ListNode {
	slow, fast := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	return slow
}

func doRemoveNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{
		Val:  0,
		Next: head,
	}
	x := findFromEnd(dummy, n)
	x.Next = x.Next.Next
	return dummy.Next
}
