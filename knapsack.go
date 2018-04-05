package main

import "fmt"

func knapsack(weight, n int, weights, values []int) int {
	dp := make([]int, weight+1)
	for i := 0; i < n; i++ {
		w, v := weights[i], values[i]
		// 因为dp中下标大的要依赖小的来计算，如果小的提前有值了，会导致大的计算错误
		// 所以j要从大往小遍历
		for j := weight; j >= 1; j-- {
			//for j := 1; j <= weight; j++ {
			if j >= w {
				if dp[j] < dp[j-w]+v {
					dp[j] = dp[j-w] + v
				}
			}
		}
	}

	return dp[weight]
}

func main() {
	fmt.Println(knapsack(5, 3, []int{1, 2, 3}, []int{6, 10, 12}))
}
