package golang

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	var ans int

	window := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		window[s[r]]++
		for window[s[r]] > 1 {
			window[s[l]]--
			l++
		}
		ans = maxInt(ans, r-l+1)
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
