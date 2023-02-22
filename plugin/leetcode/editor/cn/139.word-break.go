package golang

// leetcode submit region begin(Prohibit modification and deletion)
func wordBreak(s string, wordDict []string) bool {
	wordDictSet := make(map[string]bool, len(wordDict))
	for _, word := range wordDict {
		wordDictSet[word] = true
	}

	//	状态：dp[i] 表示字符串 s 前 i 个字符组成的字符串 s[0..i−1] 是否能被空格拆分成若干个字典中出现的单词。
	dp := make([]bool, len(s)+1)
	// 默认 j=0 时为空串
	dp[0] = true
	for i := 1; i <= len(s); i++ {
		// 枚举 s[0..i−1] 中的分割点 j
		for j := 0; j < i; j++ {
			if dp[j] && wordDictSet[s[j:i]] {
				dp[i] = true
				break
			}
		}
	}
	return dp[len(s)]
}

//leetcode submit region end(Prohibit modification and deletion)
