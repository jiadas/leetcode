package main

// 这种数在二进制表示中有且只有一个奇数位为 1，例如 16（10000）
func isPowerOfFour(num int) bool {
	return num > 0 && num&(num-1) == 0 && num&0x55555555 != 0
}
