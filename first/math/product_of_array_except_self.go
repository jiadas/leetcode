package main

import "fmt"

// Given numbers [2, 3, 4, 5], regarding the third number 4, the product of array except 4 is 2*3*5 which consists of two parts: left 2*3 and right 5.
// The product is left*right.
func productExceptSelf(nums []int) []int {
	ret := make([]int, len(nums))
	ret[0] = 1
	// left
	for i := 1; i < len(nums); i++ {
		ret[i] = ret[i-1] * nums[i-1]
	}

	right := nums[len(nums)-1]
	for i := len(nums) - 2; i >= 0; i-- {
		ret[i] *= right
		right *= nums[i]
	}
	return ret
}

func main() {
	fmt.Println(productExceptSelf([]int{1, 2, 3, 4}))
}
