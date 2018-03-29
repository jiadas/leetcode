package main

import (
	"fmt"
)

func subsets(nums []int) [][]int {
	result := [][]int{{}}
	if len(nums) == 0 {
		return result
	}

	n := len(nums)
	for i := 1; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			result = append(result, getSubset(nums[j+1:], i, []int{nums[j]}, nil)...)
		}
	}
	return result
}

func getSubset(nums []int, n int, sub []int, result [][]int) [][]int {
	if len(sub) == n {
		return append(result, append([]int(nil), sub...))
	}
	for i := 0; i < len(nums); i++ {
		result = append(result, getSubset(nums[i+1:], n, append(sub, nums[i]), nil)...)
	}
	return result
}

func main() {
	fmt.Println(subsets([]int{1, 2, 3}))
	fmt.Println(subsets([]int{1, 2, 3, 4}))
}
