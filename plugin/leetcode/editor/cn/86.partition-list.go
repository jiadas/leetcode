package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func partition(head *ListNode, x int) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	small := new(ListNode)
	smallHead := small
	large := new(ListNode)
	largeHead := large
	for head != nil {
		if head.Val >= x {
			large.Next = head
			large = large.Next
		} else {
			small.Next = head
			small = small.Next
		}
		head = head.Next
	}
	large.Next = nil
	small.Next = largeHead.Next

	return smallHead.Next
}

// leetcode submit region end(Prohibit modification and deletion)

func doPartition(nums []int, x int) []int {
	return listToIntSlice(partition(intSliceToList(nums), x))
}
