package main

import (
	"fmt"
	"sort"
)

func subsetsWithDup(nums []int) [][]int {
	result := [][]int{{}}
	if len(nums) == 0 {
		return result
	}

	sort.Ints(nums)

	n := len(nums)
	for i := 1; i <= n; i++ {
		for j := 0; j <= n-i; j++ {
			if j > 0 && nums[j] == nums[j-1] {
				continue
			}
			result = append(result, getSubsetsWithDup(nums[j+1:], i, []int{nums[j]}, nil)...)
		}
	}

	return result
}

func getSubsetsWithDup(nums []int, n int, sub []int, result [][]int) [][]int {
	if len(sub) == n {
		return append(result, append([]int(nil), sub...))
	}
	for i := 0; i < len(nums); i++ {
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}
		result = append(result, getSubsetsWithDup(nums[i+1:], n, append(sub, nums[i]), nil)...)
	}
	return result
}

func main() {
	fmt.Println(subsetsWithDup([]int{1, 2, 2}))
}
