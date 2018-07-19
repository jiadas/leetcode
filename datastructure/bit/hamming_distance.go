package main

func hammingDistance(x int, y int) int {
	z := x ^ y
	var count int
	for z != 0 {
		if z&1 == 1 {
			count++
		}
		z = z >> 1

		// 使用 z&(z-1) 去除 z 位级表示最低的那一位
		// z &= (z-1)
		// count++
	}
	return count
}
