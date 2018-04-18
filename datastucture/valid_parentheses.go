package main

import "fmt"

func isValid(s string) bool {
	stack := make([]byte, 0, len(s))
	for i := 0; i < len(s); i++ {
		v := s[i]
		if v == '(' || v == '[' || v == '{' {
			stack = append(stack, v)
			continue
		}
		if len(stack) == 0 {
			return false
		}
		var peek byte
		peek, stack = stack[len(stack)-1], stack[:len(stack)-1]
		if v == ')' && peek != '(' {
			return false
		}
		if v == ']' && peek != '[' {
			return false
		}
		if v == '}' && peek != '{' {
			return false
		}
	}
	return len(stack) == 0
}

func main() {
	fmt.Println(isValid("()[]{}"))
	fmt.Println(isValid("["))
}
