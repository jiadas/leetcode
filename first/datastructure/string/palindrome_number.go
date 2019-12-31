package main

import "fmt"

// 将整数分成左右两部分，右边那部分需要转置，然后判断这两部分是否相等
func isPalindrome(x int) bool {
	if x == 0 {
		return true
	}
	if x < 0 {
		return false
	}
	if x%10 == 0 {
		return false
	}
	var right int
	for x > right {
		right = right*10 + x%10
		x /= 10
	}
	return x == right || x == right/10
}

func main() {
	fmt.Println(isPalindrome(121))
	fmt.Println(isPalindrome(10))
}
