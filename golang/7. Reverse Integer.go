package golang

import "math"

func reverse(x int) int {
	rev := make([]int, 0, 32)
	tmp := x
	if x < 0 {
		tmp = -x
	}
	for tmp != 0 {
		rev = append(rev, tmp%10)
		tmp = tmp / 10
	}

	var result int
	for _, t := range rev {
		result = t + result*10
		if result > math.MaxInt32 {
			return 0
		}
	}
	if x < 0 {
		result = -result
	}
	return result
}

// https://leetcode.com/problems/reverse-integer/discuss/4060/My-accepted-15-lines-of-code-for-Java
func reverseFD(x int) int {
	var result int
	for x != 0 {
		tail := x % 10
		result = result*10 + tail
		if result > math.MaxInt32 || result < -math.MaxInt32 {
			return 0
		}
		x /= 10
	}
	return result
}
