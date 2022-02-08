package golang

// leetcode submit region begin(Prohibit modification and deletion)
func findAnagrams(s string, p string) []int {
	// return doFindAnagrams(s, p)
	return doFindAnagramsWithTemplate(s, p)
}

func doFindAnagramsWithTemplate(s string, p string) []int {
	needs, window := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(p); i++ {
		needs[p[i]]++
	}

	var ans []int
	var valid int
	for l, r := 0, 0; r < len(s); r++ {
		rc := s[r]
		if needs[rc] > 0 {
			window[rc]++
			if window[rc] == needs[rc] {
				valid++
			}
		}
		for r-l+1 >= len(p) {
			if valid == len(needs) {
				ans = append(ans, l)
			}
			lc := s[l]
			l++
			if needs[lc] > 0 {
				if window[lc] == needs[lc] {
					valid--
				}
				window[lc]--
			}
		}
	}

	return ans
}

func doFindAnagrams(s string, p string) []int {
	var ans []int

	lenS, lenP := len(s), len(p)
	if lenP > lenS {
		return ans
	}

	allZero := func(count []int) bool {
		for _, c := range count {
			if c != 0 {
				return false
			}
		}
		return true
	}

	count := make([]int, 26)
	for i := 0; i < lenP; i++ {
		count[p[i]-'a']++
		count[s[i]-'a']--
	}
	if allZero(count) {
		ans = append(ans, 0)
	}

	for right := lenP; right < lenS; right++ {
		count[s[right]-'a']--
		// 固定窗口大小 lenP, left=right-lenP
		count[s[right-lenP]-'a']++
		if allZero(count) {
			// new left = right-lenP+1
			ans = append(ans, right-lenP+1)
		}
	}

	return ans
}

// leetcode submit region end(Prohibit modification and deletion)
