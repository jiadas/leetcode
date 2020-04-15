package golang

import (
	"strconv"
)

// https://leetcode.com/problems/multiply-strings/discuss/17605/Easiest-JAVA-Solution-with-Graph-Explanation
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	pos := make([]int, m+n)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			d1, d2 := int(num1[i]-'0'), int(num2[j]-'0')
			mul := d1 * d2
			p1, p2 := i+j, i+j+1
			sum := mul + pos[p2]
			pos[p1] += sum / 10
			pos[p2] = sum % 10
		}
	}

	var result string
	for _, po := range pos {
		if !(len(result) == 0 && po == 0) {
			result += strconv.Itoa(po)
		}
	}
	if len(result) == 0 {
		return "0"
	}
	return result
}
