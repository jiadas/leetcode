package main

import (
	"fmt"
)

func wordBreak(s string, wordDict []string) bool {
	dp := make([]bool, len(s)+1) // dp 表示 s 中下标 i 之前的字符串是否包含字典里的单词
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		for _, word := range wordDict {
			l := len(word)
			if i >= l && word == string(s[i-l:i]) {
				dp[i] = dp[i] || dp[i-l]
			}
		}
	}
	return dp[len(s)]
}

func main() {
	fmt.Println(wordBreak("leetcode", []string{"leet", "code"}))
}
