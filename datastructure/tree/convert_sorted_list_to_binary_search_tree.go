package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func sortedListToBST(head *ListNode) *TreeNode {
	if head == nil {
		return nil
	}
	if head.Next == nil {
		return &TreeNode{Val: head.Val}
	}

	f, s := head, head
	p := s
	for f != nil && f.Next != nil {
		f = f.Next.Next
		p = s
		s = s.Next
	}
	p.Next = nil

	return &TreeNode{
		Val:   s.Val,
		Left:  sortedListToBST(head),
		Right: sortedListToBST(s.Next),
	}
}
