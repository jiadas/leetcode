package main

import "fmt"

func nextGreaterElement(findNums []int, nums []int) []int {
	m := make(map[int]int)
	stack := make([]int, len(nums))
	for _, n := range nums {
		for len(stack) != 0 && stack[len(stack)-1] < n {
			m[stack[len(stack)-1]] = n
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, n)
	}
	var ret []int
	for _, fn := range findNums {
		g, ok := m[fn]
		if !ok {
			g = -1
		}
		ret = append(ret, g)
	}
	return ret
}

func main() {
	fmt.Println(nextGreaterElement([]int{4, 1, 2}, []int{1, 3, 4, 2}))
}
