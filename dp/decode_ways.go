package main

import (
	"fmt"
	"strconv"
)

func numDecodings(s string) int {
	n := len(s)
	if n == 0 {
		return 0
	}
	dp := make([]int, n+1) // dp[i] 表示长度为 i 的数字字符串可以解码的种数
	dp[0] = 1
	if s[0] != '0' {
		dp[1] = 1
	}
	b := []byte(s)
	for i := 2; i <= n; i++ {
		one, _ := strconv.Atoi(string(b[i-1 : i]))
		if one != 0 {
			dp[i] += dp[i-1]
		}
		if b[i-2] != '0' {
			two, _ := strconv.Atoi(string(b[i-2 : i]))
			if two <= 26 {
				dp[i] += dp[i-2]
			}
		}
	}
	return dp[n]
}

func main() {
	//fmt.Println(numDecodings("12"))
	fmt.Println(numDecodings("10"))
}
