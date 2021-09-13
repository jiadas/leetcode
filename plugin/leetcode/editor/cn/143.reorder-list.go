package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reorderList(head *ListNode) {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return
	}
	mid := findMiddleNode(head)
	right := mid.Next
	mid.Next = nil

	rightDummy := new(ListNode)
	for right != nil {
		next := right.Next

		right.Next = rightDummy.Next
		rightDummy.Next = right

		right = next
	}

	right = rightDummy.Next
	left := head
	for left != nil && right != nil {
		lt, rt := left.Next, right.Next

		right.Next = left.Next
		left.Next = right

		left, right = lt, rt
	}
}

// leetcode submit region end(Prohibit modification and deletion)

func doReorderList(nums []int) []int {
	head := intSliceToList(nums)
	reorderList(head)
	return listToIntSlice(head)
}
