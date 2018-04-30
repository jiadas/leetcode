package main

import "fmt"

func nextGreaterElements(nums []int) []int {
	l := len(nums)
	ret := make([]int, l)
	stack := make([]int, l)
	for i := 0; i < l; i++ {
		for len(stack) != 0 && nums[stack[len(stack)-1]] < nums[i] {
			ret[stack[len(stack)-1]] = nums[i]
			stack = stack[:len(stack)-1]
		}
		stack = append(stack, i)
	}
	// 栈里剩下的都需要循环往回找
	for si, i := range stack {
		if si == 0 {
			ret[i] = -1
		} else {
			// 最先入栈的那个元素肯定是最大的，所以只需要遍历该元素下标之前的
			for j := 0; j <= stack[0]; j++ {
				if nums[j] > nums[i] {
					ret[i] = nums[j]
					break
				} else if j == stack[0] {
					ret[i] = -1
				}
			}
		}
	}
	return ret
}

func nextGreaterElements1(nums []int) []int {
	l := len(nums)
	ret := make([]int, l)
	for i := 0; i < l; i++ {
		ret[i] = -1
	}
	stack := make([]int, l)
	for i := 0; i < 2*l; i++ {
		n := nums[i%l]
		for len(stack) != 0 && nums[stack[len(stack)-1]] < n {
			ret[stack[len(stack)-1]] = n
			stack = stack[:len(stack)-1]
		}
		if i < l {
			stack = append(stack, i)
		}
	}
	return ret
}

func main() {
	fmt.Println(nextGreaterElements([]int{9, 8, 10, 7, 3, 2, 1, 6}))
	fmt.Println(nextGreaterElements([]int{1, 1, 1, 1, 1, 1}))
	fmt.Println(nextGreaterElements([]int{0, -2, -3}))

	fmt.Println(nextGreaterElements1([]int{9, 8, 10, 7, 3, 2, 1, 6}))
	fmt.Println(nextGreaterElements1([]int{1, 1, 1, 1, 1, 1}))
	fmt.Println(nextGreaterElements1([]int{0, -2, -3}))
}
