package main

import "fmt"

func countSubstrings(s string) int {
	var ret int
	for i := 0; i < len(s); i++ {
		ret = extendPalindrome(s, i, i, ret)   // 以一个字母为中心开始扩展
		ret = extendPalindrome(s, i, i+1, ret) // 以两个字母为中心开始扩展
	}
	return ret
}

func extendPalindrome(s string, left, right int, count int) int {
	for left >= 0 && right < len(s) && s[left] == s[right] {
		count++
		left--
		right++
	}
	return count
}

func main() {
	fmt.Println(countSubstrings("aaa"))
}
