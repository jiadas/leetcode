package main

import "fmt"

func combine(n int, k int) [][]int {
	if n == 0 || k == 0 || k > n {
		return nil
	}
	var nums []int
	for i := 1; i <= n; i++ {
		nums = append(nums, i)
	}
	return doCombine(nums, k, nil, nil)
}

func doCombine(nums []int, k int, c []int, result [][]int) [][]int {
	if k == 0 {
		return append(result, append([]int(nil), c...))
	}
	for i := 0; i < len(nums); i++ {
		r := append([]int(nil), nums[i+1:]...)
		result = append(result, doCombine(r, k-1, append(c, nums[i]), nil)...)
	}
	return result
}

func main() {
	fmt.Println(combine(4, 3))
}
