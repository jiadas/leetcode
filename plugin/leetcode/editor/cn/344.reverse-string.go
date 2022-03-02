package golang

// leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

// leetcode submit region end(Prohibit modification and deletion)
