package golang

import "math"

// https://leetcode.com/problems/string-to-integer-atoi/discuss/4654/My-simple-solution
func myAtoi(str string) int {
	if str == "" {
		return 0
	}
	var (
		sign    = 1
		base, i int
	)
	for i < len(str) && str[i] == ' ' {
		i++
	}
	if i < len(str) && (str[i] == '+' || str[i] == '-') {
		if str[i] == '-' {
			sign = -1
		}
		i++
	}
	var result int
	for ; i < len(str) && (str[i] >= '0' && str[i] <= '9'); i++ {
		base = 10*base + int(str[i]-'0')
		result = sign * base
		if result > math.MaxInt32 {
			return math.MaxInt32
		}
		if result < math.MinInt32 {
			return math.MinInt32
		}
	}

	return result
}
