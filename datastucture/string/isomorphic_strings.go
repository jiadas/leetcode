package main

import "fmt"

// 记录一个字符上次出现的位置，如果两个字符串中某个字符上次出现的位置一样，那么就属于同构。
func isIsomorphic(s string, t string) bool {
	var m1, m2 [256]int
	for i := 0; i < len(s); i++ {
		if m1[s[i]] != m2[t[i]] {
			return false
		}
		m1[s[i]] = i + 1
		m2[t[i]] = i + 1
	}
	return true
}

func main() {
	fmt.Println(isIsomorphic("egg", "foo"))
}
