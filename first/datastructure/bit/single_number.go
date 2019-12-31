package main

func singleNumber(nums []int) int {
	var ret int
	for _, n := range nums {
		ret ^= n
	}
	return ret
}
