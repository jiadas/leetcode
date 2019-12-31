package main

import "fmt"

func nthUglyNumber(n int) int {
	if n <= 6 {
		return n
	}
	dp := make([]int, n)
	dp[0] = 1
	var i2, i3, i5 int
	for i := 1; i < n; i++ {
		next2, next3, next5 := dp[i2]*2, dp[i3]*3, dp[i5]*5
		dp[i] = minInt(next2, minInt(next3, next5))
		if dp[i] == next2 {
			i2++
		}
		if dp[i] == next3 {
			i3++
		}
		if dp[i] == next5 {
			i5++
		}
	}
	return dp[n-1]
}

func minInt(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func main() {
	//fmt.Println(nthUglyNumber(1005))
	fmt.Println(nthUglyNumber(10))
}
