package golang

import "strings"

func longestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	if len(strs) == 1 {
		return strs[0]
	}

	var lcp []byte
loop:
	for {
		i := len(lcp)
		var p byte
		for j, str := range strs {
			if i > len(str)-1 {
				return string(lcp)
			}
			if j == 0 {
				p = str[i]
				continue
			}
			if p != str[i] {
				break loop
			}
		}
		lcp = append(lcp, p)
	}
	return string(lcp)
}

// https://leetcode.com/problems/longest-common-prefix/discuss/6910/Java-code-with-13-lines
func longestCommonPrefixFD(strs []string) string {
	if len(strs) == 0 {
		return ""
	}
	pre := strs[0]
	for i := 1; i < len(strs); i++ {
		for !strings.HasPrefix(strs[i], pre) {
			pre = pre[:len(pre)-1]
		}
	}
	return pre
}
