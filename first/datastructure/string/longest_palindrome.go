package main

import "fmt"

func longestPalindrome(s string) int {
	var char [128]int
	for i := 0; i < len(s); i++ {
		char[s[i]]++
	}
	var ret int
	for _, c := range char {
		// Input: "ccc"
		//Output: 1
		//Expected: 3
		ret += (c / 2) * 2 // 每个字符有偶数个可以用来构成回文字符串，即某个字符可能有7个，取6个用来构成回文字符串
	}
	if ret < len(s) {
		ret++ // 这个条件下 s 中一定有单个未使用的字符存在，可以把这个字符放到回文的最中间
	}
	return ret
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("ccc"))
}
