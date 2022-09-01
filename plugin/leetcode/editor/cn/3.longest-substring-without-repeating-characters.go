package golang

// leetcode submit region begin(Prohibit modification and deletion)
func lengthOfLongestSubstring(s string) int {
	var ans int

	window := map[byte]int{}
	for l, r := 0, 0; r < len(s); r++ {
		window[s[r]]++
		// 当 window[s[r]] 值大于 1 时，说明窗口中存在重复字符，不符合条件，就该移动 left 缩小窗口了
		for window[s[r]] > 1 {
			window[s[l]]--
			l++
		}
		// 要在收缩窗口完成后更新 ans，因为窗口收缩的 for 条件是存在重复元素，换句话说收缩完成后一定保证窗口中没有重复
		ans = maxInt(ans, r-l+1)
	}
	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
