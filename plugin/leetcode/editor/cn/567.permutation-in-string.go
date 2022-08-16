package golang

// leetcode submit region begin(Prohibit modification and deletion)
func checkInclusion(s1 string, s2 string) bool {
	needs, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		needs[s1[i]]++
	}

	var match int
	for l, r := 0, 0; r < len(s2); r++ {
		if needs[s2[r]] > 0 {
			window[s2[r]]++
			if window[s2[r]] == needs[s2[r]] {
				match++
			}
		}

		for match == len(needs) && l <= r {
			if r-l+1 == len(s1) {
				return true
			}

			if needs[s2[l]] > 0 {
				if window[s2[l]] == needs[s2[l]] {
					match--
				}
				window[s2[l]]--
			}
			l++
		}
	}

	return false
}

// leetcode submit region end(Prohibit modification and deletion)
