package main

import "fmt"

func combinationSum4(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	dp := make([]int, target+1)
	dp[0] = 1

	for i := 0; i <= target; i++ {
		for j := 0; j < len(nums); j++ {
			if i >= nums[j] {
				dp[i] += dp[i-nums[j]]
			}
		}
	}
	return dp[target]
}

func main() {
	fmt.Println(combinationSum4([]int{1, 2, 3}, 4))
}
