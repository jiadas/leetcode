package golang

// leetcode submit region begin(Prohibit modification and deletion)
func equalSubstring(s string, t string, maxCost int) int {
	var cost, ans int
	for l, r := 0, 0; r < len(s); r++ {
		cost += absInt(int(s[r]) - int(t[r]))
		for cost > maxCost {
			cost -= absInt(int(s[l]) - int(t[l]))
			l++
		}
		ans = maxInt(ans, r-l+1)
	}
	return ans
}

// func maxInt(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }
//
// func absInt(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// leetcode submit region end(Prohibit modification and deletion)
