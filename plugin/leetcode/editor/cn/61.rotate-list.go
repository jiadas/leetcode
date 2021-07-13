package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	l := 1
	tail := head
	for tail.Next != nil {
		l++
		tail = tail.Next
	}

	k = k % l
	if k == 0 {
		return head
	}

	var newHead, pre *ListNode = head, nil
	for i := 0; i < l-k; i++ {
		pre = newHead
		newHead = newHead.Next

		if i == l-k-1 {
			pre.Next = nil
			tail.Next = head
		}
	}
	return newHead
}

// leetcode submit region end(Prohibit modification and deletion)
