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
		ret += (c / 2) * 2
	}
	if ret < len(s) {
		ret++
	}
	return ret
}

func main() {
	fmt.Println(longestPalindrome("abccccdd"))
	fmt.Println(longestPalindrome("ccc"))
}
