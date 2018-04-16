package main

import "fmt"

// 平方序列：1,4,9,16,.. 间隔：3,5,7,...
//
// 间隔为等差数列，使用这个特性可以得到从 1 开始的平方序列。
func isPerfectSquare(num int) bool {
	sq := 1
	for num > 0 {
		num -= sq
		sq += 2
	}

	return num == 0
}

func main() {
	fmt.Println(isPerfectSquare(10))
	fmt.Println(isPerfectSquare(25))
}
