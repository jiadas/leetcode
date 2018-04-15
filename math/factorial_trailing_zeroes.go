package main

import "fmt"

// 编程之美 2.2
func trailingZeroes(n int) int {
	var ret int
	for n != 0 {
		ret += n / 5
		n /= 5
	}

	//for i := 1; i <= n; i++ {
	//	j := i
	//	for j%5 == 0 {
	//		ret++
	//		j = j / 5
	//	}
	//}
	return ret
}

func main() {
	fmt.Println(trailingZeroes(25))
}
