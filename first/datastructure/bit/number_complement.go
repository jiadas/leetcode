package main

// 题目要求：每个位翻转一下，但是翻转的起始位置上从最高位的1开始，前面的0不能被翻转

// 利用 x ^ 1s = ~x 的特点，可以将位级表示翻转
// 要得到 1 到 i 位为 1 的 mask，1<<i-1 即可，例如将 1<<4-1 = 00010000-1 = 00001111
func findComplement(num int) int {
	if num == 0 {
		return 1
	}
	// The given integer is guaranteed to fit within the range of a 32-bit signed integer.
	mask := 1 << 30 // 将1移到第31位(1<<30)，只移到31位，1是因为最高位是符号位；2是因为下面还要对 mask 向左移动1位，如果 num 的31位是1的话
	for num&mask == 0 {
		mask >>= 1 // 为了让 mask 最高位的1和 num 最高位的1对齐
	}
	mask = mask<<1 - 1 // 让 mask 从 num 的最高位1往右都变成1
	return num ^ mask
}
