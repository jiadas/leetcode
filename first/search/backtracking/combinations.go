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
		//return append(result, c)  // 这样写会得到重复的结果
		return append(result, append([]int(nil), c...))
	}
	//for i := 0; i < len(nums); i++ {
	for i, value := range nums {
		//r := append([]int(nil), nums[i+1:]...)
		//result = append(result, doCombine(r, k-1, append(c, nums[i]), nil)...)
		result = append(result, doCombine(nums[i+1:], k-1, append(c, value), nil)...)
	}
	return result
}

func main() {
	fmt.Println(combine(6, 4))
}
