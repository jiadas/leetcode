package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func reverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil || head.Next == nil || left == right {
		return head
	}
	i := 1
	var pre, cur *ListNode = nil, head
	var betweenHead, betweenTail, leftTail *ListNode
	for cur != nil {
		if i >= left && i <= right {
			if i == left {
				betweenTail = cur
			}
			if i == right {
				betweenHead = cur
			}

			t := cur.Next
			cur.Next = pre
			pre = cur
			cur = t
		} else {
			if i+1 == left {
				leftTail = cur
			}
			if i > right {
				if betweenTail != nil {
					betweenTail.Next = cur
				}
				break
			}
			cur = cur.Next
		}
		i++
	}
	if leftTail != nil {
		leftTail.Next = betweenHead
	}
	if left > 1 {
		return head
	}
	return betweenHead
}

// leetcode submit region end(Prohibit modification and deletion)

func doReverseBetween(nums []int, left, right int) []int {
	return listToIntSlice(reverseBetween(intSliceToList(nums), left, right))
}
