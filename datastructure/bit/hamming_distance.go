package main

func hammingDistance(x int, y int) int {
	z := x ^ y
	var count int
	for z != 0 {
		if z&1 == 1 {
			count++
		}
		z = z >> 1
	}
	return count
}
