package golang

import "math"

// https://leetcode.com/problems/divide-two-integers/discuss/13407/C++-bit-manipulations
func divide(dividend int, divisor int) int {
	if dividend == math.MinInt32 && divisor == -1 {
		return math.MaxInt32
	}
	dvd, dvs := abs(dividend), abs(divisor)
	var result int
	for dvd >= dvs {
		tmp, m := dvs, 1
		for tmp<<1 <= dvd {
			tmp <<= 1
			m <<= 1
		}
		dvd -= tmp
		result += m
	}
	return signResult(dividend, divisor, result)
}

func signResult(dividend, divisor, result int) int {
	dvdSign := 1
	if dividend < 0 {
		dvdSign = 0
	}
	divSign := 1
	if divisor < 0 {
		divSign = 0
	}

	switch dvdSign ^ divSign {
	case 0:
		return result
	case 1:
		return -result
	}
	return result
}
