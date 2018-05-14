package main

import "fmt"

func permute(nums []int) [][]int {
	return doPermute(nums, nil, nil)
}

func doPermute(nums []int, p []int, result [][]int) [][]int {
	if len(nums) == 0 {
		return append(result, p)
	}
	for i, value := range nums {
		r := append([]int{}, nums[:i]...)
		r = append(r, nums[i+1:]...)
		result = append(result, doPermute(r, append([]int(nil), append(p, value)...), nil)...)
	}

	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3}))
	fmt.Println(permute([]int{1, 2}))
	fmt.Println(permute([]int{}))
}
