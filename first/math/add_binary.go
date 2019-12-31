package main

import (
	"fmt"
	"strconv"
)

func addBinary(a string, b string) string {
	as, bs := []byte(a), []byte(b)
	carry, i, j := 0, len(a)-1, len(b)-1
	var ret string
	for i >= 0 || j >= 0 || carry == 0 {
		if i >= 0 {
			carry += int(as[i] - '0')
			i--
		}
		if j >= 0 {
			carry += int(bs[j] - '0')
			j--
		}
		ret = strconv.Itoa(carry%2) + ret
		carry /= 2
	}
	return ret
}

func main() {
	fmt.Println(addBinary("11", "1"))
}
