package main

import "fmt"

func combinationSum3(k int, n int) [][]int {
	if k == 0 || n == 0 || n > 45 {
		return nil
	}

	nums := make([]int, 0, 9)
	for i := 1; i <= 9; i++ {
		nums = append(nums, i)
	}
	return doCombinationSum3(k, n, 0, nums, nil, nil)
}

func doCombinationSum3(k, n, sum int, nums []int, r []int, result [][]int) [][]int {
	if sum == n && k == 0 {
		return append(result, append([]int(nil), r...))
	}
	if sum > n {
		return result
	}

	for i := range nums {
		if sum+nums[i] <= n {
			left := append([]int(nil), nums[i+1:]...)
			result = append(result, doCombinationSum3(k-1, n, sum+nums[i], left, append(r, nums[i]), nil)...)
		}
	}
	return result
}

func main() {
	fmt.Println(combinationSum3(3, 9))
}
