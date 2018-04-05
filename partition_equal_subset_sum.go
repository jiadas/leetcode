package main

import "fmt"

func canPartition(nums []int) bool {
	var sum int
	for _, value := range nums {
		sum += value
	}
	if sum%2 != 0 {
		return false
	}

	w := sum / 2
	dp := make([]bool, w+1) // dp[i][j] 表示前前 i 个数的和是否等于 j，用一维的 dp 数组来优化空间
	dp[0] = true
	for _, num := range nums {
		for j := w; j >= 0; j-- {
			if j >= num {
				dp[j] = dp[j] || dp[j-num]
			}
		}
	}
	return dp[w]
}

func main() {
	fmt.Println(canPartition([]int{1, 5, 11, 5}))
}
