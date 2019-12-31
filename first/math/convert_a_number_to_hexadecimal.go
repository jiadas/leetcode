package main

import "fmt"

/*
Basic idea: each time we take a look at the last four digits of
            binary verion of the input, and maps that to a hex char
            shift the input to the right by 4 bits, do it again
            until input becomes 0.

*/
func toHex(num int) string {
	m := []string{"0", "1", "2", "3", "4", "5", "6", "7", "8", "9", "a", "b", "c", "d", "e", "f"}
	if num == 0 {
		return "0"
	}
	var result string
	//  must add result.size()<8 in the while condition, or it will loop forever with a negative input
	for num != 0 && len(result) < 8 {
		result = m[num&15] + result
		num = num >> 4
	}
	return result
}

func main() {
	fmt.Println(toHex(-1))
}
