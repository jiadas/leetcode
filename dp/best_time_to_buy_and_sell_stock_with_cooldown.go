package main

import "fmt"

// 对于每一天，有三种动作，buy, sell, cooldown, sell 和 cooldown 可以合并成一种状态，因为手里最终没有股票。
// 最终需要的结果是 sell，即手里股票卖了获得最大利润。
// 我们可以用两个数组来记录当前持股和未持股的状态，令 sell[i] 表示第i天未持股时，获得的最大利润， buy[i] 表示第i天持有股票时，获得的最大利润。
// 对于 sell[i] ，最大利润有两种可能，一是今天没动作跟昨天未持股状态一样，二是今天卖了股票，所以状态转移方程如下：
//    sell[i] = max{sell[i - 1], buy[i-1] + prices[i]}
//
// 对于 buy[i] ，最大利润有两种可能，一是今天没动作跟昨天持股状态一样，二是前天卖了股票，今天买了股票，因为 cooldown 只能隔天交易，所以今天买股票要追溯到前天的状态。状态转移方程如下：
//    buy[i] = max{buy[i-1], sell[i-2] - prices[i]}
//
// 最终我们要求的结果是 sell[n - 1] ，表示最后一天结束时，手里没有股票时的最大利润。
func maxProfit(prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	sell := make([]int, n) // sell[i] 表示在第 i 天做卖出操作所获得的最大收益
	buy := make([]int, n)  // buy[i] 表示在第 i 天做买入操作所获得的最大收益
	sell[0] = 0
	buy[0] = -prices[0]
	for i := 1; i < n; i++ {
		sell[i] = sell[i-1]
		if sell[i-1] < buy[i-1]+prices[i] {
			sell[i] = buy[i-1] + prices[i]
		}

		buy[i] = buy[i-1]
		var yesterdaySell int
		if i > 1 {
			yesterdaySell = sell[i-2]
		}
		if buy[i-1] < yesterdaySell-prices[i] {
			buy[i] = yesterdaySell - prices[i]
		}
	}

	return sell[n-1]
}

func main() {
	fmt.Println(maxProfit([]int{1, 2, 3, 0, 2}))
}
