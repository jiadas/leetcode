package main

import (
	"fmt"
	"strings"
)

func findLongestWord(s string, d []string) string {
	var result string
	for _, value := range d {
		for i, j := 0, 0; i < len(s) && j < len(value); i++ {
			if s[i] == value[j] {
				j++
			}
			if j == len(value) {
				if len(value) > len(result) || (len(value) == len(result) && strings.Compare(value, result) < 0) {
					result = value
				}
			}
		}
	}
	return result
}

func main() {
	fmt.Println(findLongestWord("abpcplea", []string{"ale", "apple", "monkey", "plea"}))
}
