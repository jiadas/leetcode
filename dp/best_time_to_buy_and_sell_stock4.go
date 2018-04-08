package main

import "fmt"

func maxProfit4(k int, prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	// 如果 price 的对数([买，卖])比交易次数小，就可以全程低买高卖实现利益最大化
	if k >= n>>1 {
		var maxP int
		for i := 1; i < n; i++ {
			// 低买高卖
			if prices[i] > prices[i-1] {
				maxP += prices[i] - prices[i-1]
			}
		}
		return maxP
	}

	/**
	 * dp[i, j] represents the max profit up until prices[j] using at most i transactions.
	 * dp[i, j] = max(dp[i, j-1], prices[j] - prices[jj] + dp[i-1, jj]) { jj in range of [0, j-1] }
	 *          = max(dp[i, j-1], prices[j] + max(dp[i-1, jj] - prices[jj]))
	 * dp[0, j] = 0; 0 transactions makes 0 profit
	 * dp[i, 0] = 0; if there is only one price data point you can't make any transaction.
	 */
	var dp [][]int // dp[i][j] 表示到 price[j] 为止用 i 次交易获取的最大利益
	for i := 0; i <= k; i++ {
		dp = append(dp, make([]int, n))
	}
	for i := 1; i <= k; i++ {
		localMax := dp[i-1][0] - prices[0]
		for j := 1; j < n; j++ {
			dp[i][j] = max(dp[i][j-1], prices[j]+localMax)
			localMax = max(localMax, dp[i-1][j]-prices[j])
		}
	}

	return dp[k][n-1]
}

func max(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit4(4, []int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(maxProfit4(2, []int{3, 3, 5, 0, 0, 3, 1, 4}))
}
