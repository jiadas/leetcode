package main

// 题目要求：每个位翻转一下，但是翻转的起始位置上从最高位的1开始的，前面的0不能被翻转

// 利用 x ^ 1s = ~x 的特点，可以将位级表示翻转
// 要得到 1 到 i 位为 1 的 mask，1<<i-1 即可，例如将 1<<4-1 = 00010000-1 = 00001111
func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	mask := 1 << 30 // 右移多少个 1，即 1 后面有多少个 0
	for num&mask == 0 {
		mask >>= 1
	}
	mask = mask<<1 - 1
	return num ^ mask
}
