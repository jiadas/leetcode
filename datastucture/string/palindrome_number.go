package main

import "fmt"

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
