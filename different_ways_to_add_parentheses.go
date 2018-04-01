package main

import (
	"fmt"
	"strconv"
)

func diffWaysToCompute(input string) []int {
	var result []int
	for i, b := range input {
		if b == '+' || b == '-' || b == '*' {
			left := diffWaysToCompute(input[:i])
			right := diffWaysToCompute(input[i+1:])
			for _, l := range left {
				for _, r := range right {
					switch b {
					case '+':
						result = append(result, l+r)
					case '-':
						result = append(result, l-r)
					case '*':
						result = append(result, l*r)
					}
				}
			}
		}
	}
	if len(result) == 0 {
		v, _ := strconv.Atoi(input)
		result = append(result, v)
	}
	return result
}

func main() {
	fmt.Println(diffWaysToCompute("2-1-1"))
	fmt.Println(diffWaysToCompute("2*3-4*5"))
}
