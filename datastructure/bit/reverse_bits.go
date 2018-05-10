package main

import "fmt"

func reverseBits(n int) int {
	var ret int
	for i := 0; i < 32; i++ {
		ret <<= 1
		ret |= n & 1
		n >>= 1
	}
	return ret
}

// for 8 bit binary number abcdefgh, the process is as follow:
// abcdefgh -> efghabcd -> ghefcdab -> hgfedcba
func reverseBitsO1(n int) int {
	n = (n >> 16) | (n << 16)
	n = ((n & 0xff00ff00) >> 8) | ((n & 0x00ff00ff) << 8)
	n = ((n & 0xf0f0f0f0) >> 4) | ((n & 0x0f0f0f0f) << 4)
	n = ((n & 0xcccccccc) >> 2) | ((n & 0x33333333) << 2)
	n = ((n & 0xaaaaaaaa) >> 1) | ((n & 0x55555555) << 1)
	return n
}

func main() {
	fmt.Println(reverseBits(43261596))
	fmt.Println(reverseBitsO1(43261596))
}
