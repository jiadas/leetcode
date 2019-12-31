package main

import "fmt"

func isAnagram(s string, t string) bool {
	var char [26]int
	for i := 0; i < len(s); i++ {
		char[s[i]-'a']++
	}
	for i := 0; i < len(t); i++ {
		char[t[i]-'a']--
	}
	for _, value := range char {
		if value != 0 {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(isAnagram("far", "var"))
}
