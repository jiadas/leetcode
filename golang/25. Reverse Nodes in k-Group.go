package golang

// 递归：https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/11423/Short-but-recursive-Java-code-with-comments
// 递归虽然简洁，但是不符合题目里要求的常数空间复杂度
// 所以还是得遍历：https://leetcode.com/problems/reverse-nodes-in-k-group/discuss/11440/Non-recursive-Java-solution-and-idea
func reverseKGroup(head *ListNode, k int) *ListNode {
	if head == nil || k == 1 {
		return head
	}
	dummy := &ListNode{Next: head}
	begin := dummy
	var i int
	for head != nil {
		i++
		if i%k == 0 {
			begin = reverseList(begin, head.Next)
			head = begin.Next
		} else {
			head = head.Next
		}
	}
	return dummy.Next
}

func reverseList(begin, end *ListNode) *ListNode {
	first, curr := begin.Next, begin.Next
	pre := begin
	var next *ListNode
	for curr != end {
		next = curr.Next
		curr.Next = pre
		pre = curr
		curr = next
	}
	begin.Next = pre
	first.Next = end
	return first
}
