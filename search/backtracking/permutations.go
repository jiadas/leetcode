package main

import "fmt"

func permute(nums []int) [][]int {
	return doPermute(nums, nil, nil)
}

func doPermute(nums []int, p []int, result [][]int) [][]int {
	if len(nums) == 0 {
		//return append(result, p) // 这样写结果也会被覆盖，出现重复
		return append(result, append([]int(nil), p...))
	}
	for i, value := range nums {
		r := append([]int{}, nums[:i]...)
		r = append(r, nums[i+1:]...)
		//result = append(result, doPermute(r, append([]int(nil), append(p, value)...), nil)...)
		result = append(result, doPermute(r, append(p, value), nil)...)
	}

	return result
}

func main() {
	fmt.Println(permute([]int{1, 2, 3, 4, 5, 6, 7}))
	fmt.Println(permute([]int{1, 2}))
	fmt.Println(permute([]int{}))
}
