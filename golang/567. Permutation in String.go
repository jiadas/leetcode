package golang

import "math"

func checkInclusion(s1 string, s2 string) bool {
	return doCheckInclusion2(s1, s2)
}

// 对 "76. Minimum Window Substring" 解法的借鉴
func doCheckInclusion1(s1, s2 string) bool {
	tm, cm := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(s1); i++ {
		tm[s1[i]]++
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
	for l, r := 0, 0; r < len(s2); r++ {
		if tm[s2[r]] > 0 {
			cm[s2[r]]++
		}
		for check() && l <= r {
			if r-l+1 < minLen {
				minLen = r - l + 1
			}
			if _, ok := tm[s2[l]]; ok {
				cm[s2[l]]--
			}
			l++
		}
	}
	if minLen == len(s1) {
		return true
	}
	return false
}

// https://leetcode.com/problems/permutation-in-string/discuss/102588/Java-Solution-Sliding-Window
func doCheckInclusion2(s1, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}

	count := make([]int, 26)
	for i := range s1 {
		count[s1[i]-'a']++
		count[s2[i]-'a']--
	}
	if allZero(count) {
		return true
	}

	for i := len1; i < len2; i++ {
		count[s2[i]-'a']--
		count[s2[i-len1]-'a']++
		if allZero(count) {
			return true
		}
	}

	return false
}

func allZero(count []int) bool {
	for _, c := range count {
		if c != 0 {
			return false
		}
	}
	return true
}
