package main

import (
	"fmt"
)

func addStrings(num1 string, num2 string) string {
	nb1, nb2 := []byte(num1), []byte(num2)
	carry, i, j := 0, len(num1)-1, len(num2)-1
	var ret string
	for i >= 0 || j >= 0 || carry > 0 {
		if i >= 0 {
			carry += int(nb1[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(nb2[j] - '0')
			j--
		}
		ret = string(byte(carry%10)+'0') + ret
		//ret = strconv.Itoa(carry%10) + ret
		carry /= 10
	}
	return ret
}

func main() {
	fmt.Println(addStrings("123", "949"))
}
