package golang

import "strings"

func isPalindrome125(s string) bool {
	if s == "" {
		return true
	}
	s = strings.ToLower(s)
	i, j := 0, len(s)-1
	for i < j {
		for !isLetterOrNum(s[i]) && i < j {
			i++
		}
		for !isLetterOrNum(s[j]) && i < j {
			j--
		}
		if i < j {
			if s[i] != s[j] {
				return false
			}
			i++
			j--
		}
	}
	return true
}

func isLetterOrNum(b byte) bool {
	if b >= 'a' && b <= 'z' || b >= 'A' && b <= 'Z' || b >= '0' && b <= '9' {
		return true
	}
	return false
}
