package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func oddEvenList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	oddTail, evenHead := head, head.Next
	odd, even := head, head.Next
	for odd != nil && even != nil {
		odd.Next = even.Next
		if even.Next != nil {
			even.Next = even.Next.Next
		}

		even = even.Next
		odd = odd.Next
		if odd != nil {
			oddTail = odd
		}
	}
	oddTail.Next = evenHead
	return head
}

// leetcode submit region end(Prohibit modification and deletion)

func doOddEvenList(nums []int) []int {
	return listToIntSlice(oddEvenList(intSliceToList(nums)))
}
