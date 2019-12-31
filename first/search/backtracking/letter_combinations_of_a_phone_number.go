package main

import "fmt"

var keys = []string{"", "", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	return combination("", digits, 0, nil)
}

func combination(r string, digits string, offset int, result []string) []string {
	if offset == len(digits) {
		return append(result, r)
	}

	index := int(digits[offset] - '0')
	for _, v := range keys[index] {
		result = append(result, combination(r+string(v), digits, offset+1, nil)...)
	}
	return result
}

func main() {
	fmt.Println(letterCombinations("23"))
	fmt.Println(letterCombinations(""))
}
