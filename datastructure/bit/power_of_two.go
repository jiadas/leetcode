package main

// 二进制表示只有一个 1 存在
// 利用 1000 & 0111 == 0 这种性质
func isPowerOfTwo(n int) bool {
	return n > 0 && (n&(n-1)) == 0
}
