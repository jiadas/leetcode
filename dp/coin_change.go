package main

import "fmt"

func coinChange(coins []int, amount int) int {
	dp := make([]int, amount+1) // dp[j] 表示金额为 j 时，最少的硬币数
	for i := 1; i <= amount; i++ {
		dp[i] = amount + 1
	}

	for i := 0; i <= amount; i++ {
		for _, coin := range coins {
			if i >= coin {
				if dp[i-coin]+1 < dp[i] {
					dp[i] = dp[i-coin] + 1
				}
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

func main() {
	fmt.Println(coinChange([]int{1, 2, 5}, 11))
}
