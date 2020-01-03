package golang

// https://leetcode.com/problems/longest-palindromic-substring/discuss/2928/Very-simple-clean-java-solution
func longestPalindrome(s string) string {
	if len(s) < 2 {
		return s
	}
	var start, maxLen int
	for i := range s {
		s1, l1 := extendPalindrome(s, i, i)
		s2, l2 := extendPalindrome(s, i, i+1)
		if l1 < l2 {
			s1 = s2
			l1 = l2
		}
		if l1 > maxLen {
			start = s1
			maxLen = l1
		}
	}
	return s[start : start+maxLen]
}

func extendPalindrome(s string, i, j int) (start, length int) {
	for i >= 0 && j < len(s) && s[i] == s[j] {
		i--
		j++
	}
	start = i + 1
	length = j - i - 1
	return
}
