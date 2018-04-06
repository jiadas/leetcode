package main

import "fmt"

func climbStairs(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var result int
	preStep1, preStep2 := 2, 1
	for i := 3; i <= n; i++ {
		result = preStep1 + preStep2
		preStep2 = preStep1
		preStep1 = result
	}
	return result
}

// 递归超时
func climbStairsOT(n int) int {
	if n == 0 {
		return 0
	}
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	var result int
	for _, value := range []int{1, 2} {
		result += climbStairs(n - value)
	}
	return result
}

func main() {
	fmt.Println(climbStairs(3))
	fmt.Println(climbStairs(44))
}
