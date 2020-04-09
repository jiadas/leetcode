package golang

import "fmt"

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}
	result := "1"
	for i := 1; i < n; i++ {
		result = build38(result)
	}
	return result
}

func build38(s string) string {
	i := 0
	var c byte
	var result string
	for i < len(s) {
		c = s[i]
		count := 0
		for i < len(s) && c == s[i] {
			i++
			count++
		}
		result += fmt.Sprintf("%d", count)
		result += string(c)
	}
	return result
}
