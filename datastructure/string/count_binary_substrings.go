package main

import "fmt"

// First, I count the number of 1 or 0 grouped consecutively.
// For example “0110001111” will be [1, 2, 3, 4].
//
// Second, for any possible substrings with 1 and 0 grouped consecutively, the number of valid substring will be the minimum number of 0 and 1.
// For example “0001111”, will be min(3, 4) = 3, ("01", "0011", "000111")
func countBinarySubstrings(s string) int {
	var ret int
	pre, cur := 0, 1
	for i := 1; i < len(s); i++ {
		if s[i] == s[i-1] {
			cur++
		} else {
			pre = cur
			cur = 1
		}

		if pre >= cur {
			ret++
		}
	}
	return ret
}

func main() {
	fmt.Println(countBinarySubstrings("00110011"))
	fmt.Println(countBinarySubstrings("10101"))
}
