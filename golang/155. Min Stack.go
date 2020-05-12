package golang

import "math"

// 自己写的代码，虽然也想到了用一个栈来存历史的最小值，但是没有用重复元素跟原栈保持数量上的一致，
// 导致在出栈的时候要写好多代码来维持最小栈的状态
// type MinStack struct {
// 	min  int
// 	mins []int
// 	data []int
// }
//
// /** initialize your data structure here. */
// func Constructor() MinStack {
// 	return MinStack{min: math.MaxInt64}
// }
//
// func (s *MinStack) Push(x int) {
// 	s.data = append(s.data, x)
// 	if s.min >= x {
// 		s.min = x
// 		s.mins = append(s.mins, x)
// 	}
// }
//
// func (s *MinStack) Pop() {
// 	top := s.Top()
// 	if top == s.min {
// 		if len(s.mins) > 0 {
// 			s.mins = s.mins[:len(s.mins)-1]
// 		}
// 		if len(s.mins) > 0 {
// 			s.min = s.mins[len(s.mins)-1]
// 		} else {
// 			s.min = math.MaxInt64
// 		}
// 	}
//
// 	if len(s.data) > 0 {
// 		s.data = s.data[:len(s.data)-1]
// 	}
// }
//
// func (s *MinStack) Top() int {
// 	if len(s.data) > 0 {
// 		return s.data[len(s.data)-1]
// 	}
// 	return 0
// }
//
// func (s *MinStack) GetMin() int {
// 	return s.min
// }

// https://leetcode-cn.com/problems/min-stack/solution/zui-xiao-zhan-by-leetcode-solution/
type MinStack struct {
	stack    []int
	minStack []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		minStack: []int{math.MaxInt64},
	}
}

func (s *MinStack) Push(x int) {
	s.stack = append(s.stack, x)
	s.minStack = append(s.minStack, minInt(s.GetMin(), x))
}

func (s *MinStack) Pop() {
	s.stack = s.stack[:len(s.stack)-1]
	s.minStack = s.minStack[:len(s.minStack)-1]
}

func (s *MinStack) Top() int {
	return s.stack[len(s.stack)-1]
}

func (s *MinStack) GetMin() int {
	return s.minStack[len(s.minStack)-1]
}
