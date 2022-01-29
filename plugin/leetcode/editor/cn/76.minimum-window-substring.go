package golang

import "math"

// leetcode submit region begin(Prohibit modification and deletion)
func minWindow(s string, t string) string {
	needs, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		needs[t[i]]++
	}

	var valid int
	minLen := math.MaxInt64
	ansStart := -1
	for l, r := 0, 0; r < len(s); r++ {
		if needs[s[r]] > 0 {
			window[s[r]]++
			if window[s[r]] == needs[s[r]] {
				valid++
			}
		}

		for valid == len(needs) && l <= r {
			if r-l+1 < minLen {
				minLen = r - l + 1
				ansStart = l
			}
			if _, ok := needs[s[l]]; ok {
				if window[s[l]] == needs[s[l]] {
					valid--
				}
				window[s[l]]--
			}
			l++
		}
	}
	if ansStart == -1 {
		return ""
	}
	return s[ansStart : ansStart+minLen]
}

// leetcode submit region end(Prohibit modification and deletion)
