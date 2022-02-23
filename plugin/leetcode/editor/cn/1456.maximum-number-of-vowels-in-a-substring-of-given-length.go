package golang

// leetcode submit region begin(Prohibit modification and deletion)

var vowel = map[byte]struct{}{
	'a': {},
	'e': {},
	'i': {},
	'o': {},
	'u': {},
}

func maxVowels(s string, k int) int {
	var ans, count int
	for right := 0; right < len(s); right++ {
		if _, ok := vowel[s[right]]; ok {
			count++
		}
		if right >= k-1 {
			ans = maxInt(ans, count)

			if _, ok := vowel[s[right-k+1]]; ok {
				count--
			}
		}
	}
	return ans
}

// func maxInt(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// leetcode submit region end(Prohibit modification and deletion)
