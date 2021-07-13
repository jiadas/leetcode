package golang

// leetcode submit region begin(Prohibit modification and deletion)
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var ans, tail *ListNode
	pre, cur := head, head.Next
	if pre.Val != cur.Val {
		ans = pre
		tail = pre
	}
	for cur.Next != nil {
		if pre.Val != cur.Val && cur.Val != cur.Next.Val {
			if ans == nil {
				ans = cur
			}
			if tail != nil {
				tail.Next = cur
			}
			tail = cur
		}
		pre = cur
		cur = cur.Next
	}
	// check the last node
	if cur.Val != pre.Val {
		if ans == nil {
			ans = cur
		}
		if tail != nil {
			tail.Next = cur
		}
		tail = cur
	}
	if tail != nil {
		tail.Next = nil
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)

func doDeleteDuplicates(nums []int) []int {
	return listToIntSlice(deleteDuplicates(intSliceToList(nums)))
}
