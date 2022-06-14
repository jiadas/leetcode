package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteNode237(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}

// leetcode submit region end(Prohibit modification and deletion)
