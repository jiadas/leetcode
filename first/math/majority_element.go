package main

import (
	"fmt"
	"sort"
)

// 先对数组排序，最中间那个数出现次数一定多于 n / 2
func majorityElement(nums []int) int {
	sort.Ints(nums)
	return nums[len(nums)/2]
}

// 数组中有一个数字出现的次数超过数组长度的一半，也就是说它出现的次数比其他所有数字出现次数的总和还要多
func majorityElement2(nums []int) int {
	var result, times int
	for _, v := range nums {
		if times == 0 {
			result = v
			times++
		} else if v == result {
			times++
		} else {
			times--
		}
	}
	return result
}

func main() {
	fmt.Println(majorityElement([]int{1, 2, 45, 34, 3, 3, 3, 3, 3, 3, 3}))
	fmt.Println(majorityElement2([]int{1, 2, 45, 34, 3, 3, 3, 3, 3, 3, 3}))
}
