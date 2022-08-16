package golang

// leetcode submit region begin(Prohibit modification and deletion)
func reverseString(s []byte) {
	for left, right := 0, len(s)-1; left < right; left++ {
		s[left], s[right] = s[right], s[left]
		right--
	}
}

func doReverseString(s []byte) {
	// https://github.com/golang/go/wiki/SliceTricks#reversing
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		s[left], s[right] = s[right], s[left]
	}
}

// leetcode submit region end(Prohibit modification and deletion)
