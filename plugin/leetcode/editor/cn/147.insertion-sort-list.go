package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func insertionSortList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummy := new(ListNode)
	cur := head
	for cur != nil {
		next := cur.Next

		pre, i := dummy, dummy.Next
		for i != nil && i.Val <= cur.Val {
			pre = i
			i = i.Next
		}
		cur.Next = i
		pre.Next = cur

		cur = next
	}
	return dummy.Next
}

// leetcode submit region end(Prohibit modification and deletion)

func doInsertionSortList(nums []int) []int {
	return listToIntSlice(insertionSortList(intSliceToList(nums)))
}
