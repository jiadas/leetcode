package main

import "math"

func judgeSquareSum(c int) bool {
	i, j := 0, int(math.Sqrt(float64(c)))
	for i <= j {
		sum := i*i + j*j
		if sum == c {
			return true
		} else if sum > c {
			j -= 1
		} else {
			i += 1
		}
	}
	return false
}
