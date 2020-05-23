package golang

import "math"

// https://leetcode-cn.com/problems/minimum-window-substring/solution/zui-xiao-fu-gai-zi-chuan-by-leetcode-solution/
// 还有优化空间
func minWindow(s string, t string) string {
	tm, cm := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		tm[t[i]]++
	}
	check := func() bool {
		for b, i := range tm {
			if cm[b] < i {
				return false
			}
		}
		return true
	}
	minLen := math.MaxInt64
	low, high := -1, -1
	for l, r := 0, 0; r < len(s); r++ {
		if tm[s[r]] > 0 {
			cm[s[r]]++
		}
		for check() && l <= r {
			if r-l+1 < minLen {
				minLen = r - l + 1
				low, high = l, l+minLen
			}
			if _, ok := tm[s[l]]; ok {
				cm[s[l]]--
			}
			l++
		}
	}
	if low == -1 {
		return ""
	}
	return s[low:high]
}
