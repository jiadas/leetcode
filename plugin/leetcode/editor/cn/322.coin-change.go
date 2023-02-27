package golang

// leetcode submit region begin(Prohibit modification and deletion)
func coinChange(coins []int, amount int) int {
	// 状态: dp[i] 表示金额为 i 时，组成的最小硬币个数
	dp := make([]int, amount+1)
	// 初始化为最大值 dp[i]=amount+1
	for i := range dp {
		dp[i] = amount + 1
	}

	dp[0] = 0
	for i := 1; i < len(dp); i++ {
		// 最大值不能初始化成 math.MaxInt，具体原因可调试 wa1 的单测
		//dp[i] = math.MaxInt
		for j := 0; j < len(coins); j++ {
			if i-coins[j] >= 0 {
				dp[i] = minInt(dp[i], dp[i-coins[j]]+1)
			}
		}
	}

	if dp[amount] > amount {
		return -1
	}
	return dp[amount]
}

//leetcode submit region end(Prohibit modification and deletion)
