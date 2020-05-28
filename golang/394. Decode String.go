package golang

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	var stack []string
	ptr := 0
	for ptr < len(s) {
		cur := s[ptr]
		if '0' <= cur && cur <= '9' {
			// 这里没有像官方题解一样封成方法，是因为不想用 ptr 的指针来隐含传递下标的递增
			var digits string
			for ; '0' <= s[ptr] && s[ptr] <= '9'; ptr++ {
				digits += string(s[ptr])
			}
			stack = append(stack, digits)
		} else if (cur >= 'a' && cur <= 'z' || cur >= 'A' && cur <= 'Z') || cur == '[' {
			stack = append(stack, string(cur))
			ptr++
		} else {
			ptr++
			var sub []string
			for stack[len(stack)-1] != "[" {
				sub = append(sub, stack[len(stack)-1])
				stack = stack[:len(stack)-1]
			}
			stack = stack[:len(stack)-1]
			for i := len(sub)/2 - 1; i >= 0; i-- {
				opp := len(sub) - i - 1
				sub[i], sub[opp] = sub[opp], sub[i]
			}
			repTime, _ := strconv.Atoi(stack[len(stack)-1])
			stack = stack[:len(stack)-1]
			stack = append(stack, strings.Repeat(strings.Join(sub, ""), repTime))
		}
	}
	return strings.Join(stack, "")
}

// 最初的思路和官方题解一相同，最后又着了递归的道儿，彻底写不出来了
//func decodeString(s string) string {
//	var ans string
//	stack := make([]byte, 0, len(s))
//	for i := 0; i < len(s); i++ {
//		if s[i] != ']' {
//			stack = append(stack, s[i])
//		} else {
//			var es string
//			for j := len(stack) - 1; j > 0; j-- {
//				if stack[j] == '[' {
//					es = string(stack[j+1:])
//					n, left := getNum(string(stack[:j]))
//					for q := 0; q < n; q++ {
//						ans += es
//					}
//					if left != "" && s[i+1:] != "" {
//						return decodeString(left + ans + s[i+1:])
//					}
//				}
//			}
//		}
//	}
//	return ans
//}
//
//func getNum(s string) (int, string) {
//	var (
//		k    int
//		left string
//	)
//	for i := len(s) - 1; i >= 0; i-- {
//		if '0' < s[i] && s[i] < '9' {
//			k = k*10 + int(s[i]-'0')
//		} else {
//			left = s[:i+1]
//			break
//		}
//	}
//	return k, left
//}
