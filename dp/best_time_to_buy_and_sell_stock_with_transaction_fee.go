package main

import "fmt"

func maxProfitFee(prices []int, fee int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	sell := make([]int, n) // sell[i] 表示在第 i 天做卖出操作所获得的最大收益
	buy := make([]int, n)  // buy[i] 表示在第 i 天做买入操作所获得的最大收益

	sell[0] = 0
	buy[0] = -prices[0]

	for i := 1; i < n; i++ {
		sell[i] = maxP(sell[i-1], buy[i-1]+prices[i]-fee)

		buy[i] = maxP(buy[i-1], sell[i-1]-prices[i])
	}

	return sell[n-1]
}

func maxP(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfitFee([]int{1, 3, 2, 8, 4, 9}, 2))
}
