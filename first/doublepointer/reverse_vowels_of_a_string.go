package main

import (
	"fmt"
	"strings"
)

func reverseVowels(s string) string {
	if len(s) == 0 {
		return s
	}

	v := []byte{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	i, j := 0, len(s)-1
	sl := make([]byte, len(s))
	// 一定要加上等号：hello
	for i <= j {
		ci := s[i]
		cj := s[j]
		if !strings.Contains(string(v), string(ci)) {
			sl[i] = ci
			i++
		} else if !strings.Contains(string(v), string(cj)) {
			sl[j] = cj
			j--
		} else {
			sl[i] = cj
			sl[j] = ci
			i++
			j--
		}
	}
	return string(sl)
}

func main() {
	fmt.Println(reverseVowels("hello"))
}
