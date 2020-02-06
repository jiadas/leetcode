package golang

// 题解：https://leetcode-cn.com/problems/swap-nodes-in-pairs/solution/liang-liang-jiao-huan-lian-biao-zhong-de-jie-di-19/
func swapPairs(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	nh := head.Next
	head.Next = swapPairs(nh.Next)
	nh.Next = head
	return nh
}
