package golang

// ListNode is a singly-linked list node.
type ListNode struct {
	Val  int
	Next *ListNode
}

func listToIntSlice(head *ListNode) []int {
	if head == nil {
		return nil
	}
	var result []int
	cur := head
	for cur != nil {
		result = append(result, cur.Val)
		cur = cur.Next
	}
	return result
}

func intSliceToList(nums []int) *ListNode {
	preHead := new(ListNode)
	var tail *ListNode
	for _, num := range nums {
		node := &ListNode{
			Val: num,
		}
		if preHead.Next == nil {
			preHead.Next = node
		}
		if tail != nil {
			tail.Next = node
		}
		tail = node
	}
	return preHead.Next
}

func roundConvertForTest(nums []int) []int {
	return listToIntSlice(intSliceToList(nums))
}
