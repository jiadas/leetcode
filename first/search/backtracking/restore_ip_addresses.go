package main

import (
	"fmt"
	"strconv"
)

func restoreIpAddresses(s string) []string {
	return doRestore(0, "", s, nil)
}

func doRestore(k int, ip string, s string, result []string) []string {
	if k == 4 || len(s) == 0 {
		if k == 4 && len(s) == 0 {
			return append(result, ip)
		}
		return result
	}
	// i不能比给定的字符串长，防止越界
	for i := 0; i <= 2 && i < len(s); i++ {
		// 位数超过1的，不能以0开头
		if i > 0 && s[0] == '0' {
			// 直接退出for循环，以为之后的组合都是非法的
			break
		}
		section := s[:i+1]
		value, _ := strconv.Atoi(section)
		if value <= 255 {
			var nextIP string
			if ip == "" {
				nextIP = section
			} else {
				nextIP = ip + "." + section
			}
			result = append(result, doRestore(k+1, nextIP, s[i+1:], nil)...)
		}
	}
	return result
}

func main() {
	fmt.Println(restoreIpAddresses("25525511135"))
}
