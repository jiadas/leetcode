package main

import (
	"fmt"
	"sort"
)

func permuteUnique(nums []int) [][]int {
	sort.Ints(nums)
	return doUniquePermute(nums, nil, nil)
}

func doUniquePermute(nums []int, p []int, result [][]int) [][]int {
	if len(nums) == 0 {
		return append(result, p)
	}

	for i, value := range nums {
		if i != 0 && nums[i-1] == nums[i] {
			continue
		}
		r := append([]int{}, nums[:i]...)
		r = append(r, nums[i+1:]...)
		// 一定要用p的copy去递归调用
		// 因为当slice的cap>len时，append操作的是同一个底层数组，元素以最终的结果为准，
		// 只有当len>cap时才会重新分配数组，这样再做append就不会影响老的数组元素
		// append([]int(nil), append(p, value)...) 等价于
		// b = make([]int, len(append(p, value)))
		// copy(b, p)
		result = append(result, doUniquePermute(r, append([]int(nil), append(p, value)...), nil)...)
	}

	return result
}

func main() {
	//fmt.Println(permuteUnique([]int{1, 2, 2}))
	//fmt.Println(permuteUnique([]int{1, 1, 2}))
	fmt.Println(permuteUnique([]int{-1, 2, -1, 2, 1, -1, 2, 1}))
}
