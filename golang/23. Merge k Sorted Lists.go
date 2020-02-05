package golang

import "container/heap"

// 中文 leetcode 官方题解：
// https://leetcode-cn.com/problems/merge-k-sorted-lists/solution/he-bing-kge-pai-xu-lian-biao-by-leetcode/
// 分治法：https://leetcode.com/problems/merge-k-sorted-lists/discuss/10522/My-simple-java-Solution-use-recursion
// 优先队列：https://leetcode.com/problems/merge-k-sorted-lists/discuss/10528/A-java-solution-based-on-Priority-Queue
func mergeKLists(lists []*ListNode) *ListNode {
	var pq PriorityQueue
	heap.Init(&pq)

	for _, list := range lists {
		if list != nil {
			heap.Push(&pq, list)
		}
	}
	head := &ListNode{}
	tail := head
	var node *ListNode
	for len(pq) > 0 {
		node = heap.Pop(&pq).(*ListNode)
		tail.Next = node
		tail = tail.Next

		node = node.Next
		if node != nil {
			heap.Push(&pq, node)
		}
	}
	return head.Next
}

type PriorityQueue []*ListNode

func (pq PriorityQueue) Len() int {
	return len(pq)
}

func (pq PriorityQueue) Less(i, j int) bool {
	return pq[i].Val < pq[j].Val
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	*pq = append(*pq, x.(*ListNode))
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil
	*pq = old[:n-1]
	return item
}
