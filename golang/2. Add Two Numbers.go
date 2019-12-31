package golang

import "fmt"

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func NewList(vals ...int) *ListNode {
	var head, tail *ListNode
	for i, val := range vals {
		node := &ListNode{
			Val:  val,
			Next: nil,
		}
		if i == 0 {
			head = node
		}
		if tail != nil {
			tail.Next = node
		}
		tail = node
	}
	return head
}

func (l *ListNode) String() string {
	var vals []int
	w := l
	for w != nil {
		vals = append(vals, w.Val)
		w = w.Next
	}
	return fmt.Sprint(vals)
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var head, tail *ListNode
	w1, w2 := l1, l2
	var o int
	for w1 != nil && w2 != nil {
		sum := w1.Val + w2.Val + o
		o = sum / 10
		node := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		if head == nil {
			head = node
		}

		if tail != nil {
			tail.Next = node
		}
		tail = node

		w1 = w1.Next
		w2 = w2.Next
	}

	left := w1
	if w2 != nil {
		left = w2
	}
	for left != nil {
		sum := left.Val + o
		o = sum / 10
		node := &ListNode{
			Val:  sum % 10,
			Next: nil,
		}

		if tail != nil {
			tail.Next = node
		}
		tail = node

		left = left.Next
		if o == 0 {
			tail.Next = left
			break
		}
	}
	if o != 0 && tail != nil {
		tail.Next = &ListNode{
			Val:  o,
			Next: nil,
		}
	}

	return head
}

// https://leetcode.com/problems/add-two-numbers/discuss/1010/Is-this-Algorithm-optimal-or-what
func addTwoNumbersFromDiscuss(l1 *ListNode, l2 *ListNode) *ListNode {
	c1, c2 := l1, l2
	sentinel := new(ListNode)
	d := sentinel
	var sum int
	for c1 != nil || c2 != nil {
		sum /= 10
		if c1 != nil {
			sum += c1.Val
			c1 = c1.Next
		}
		if c2 != nil {
			sum += c2.Val
			c2 = c2.Next
		}
		d.Next = &ListNode{
			Val:  sum % 10,
			Next: nil,
		}
		d = d.Next
	}
	if sum/10 == 1 {
		d.Next = &ListNode{
			Val:  1,
			Next: nil,
		}
	}
	return sentinel.Next
}
