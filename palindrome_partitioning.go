package main

import "fmt"

func partition(s string) [][]string {
	return doPartition(s, nil, nil)
}

func doPartition(s string, t []string, result [][]string) [][]string {
	if len(s) == 0 {
		return append(result, append([]string(nil), t...))
	}
	n := len(s)
	for i := 1; i <= n; i++ {
		sub := s[:i]
		if isPalindrome(sub) {
			result = append(result, doPartition(s[i:], append(t, sub), nil)...)
		}
	}
	return result
}

func isPalindrome(s string) bool {
	for left, right := 0, len(s)-1; left < right; left, right = left+1, right-1 {
		if s[left] != s[right] {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(partition("aab"))
}
