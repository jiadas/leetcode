package main

import (
	"fmt"
	"sort"
)

// 这是个典型的相遇问题，移动距离最小的方式是所有元素都移动到中位数。理由如下：
//
// 设 m 为中位数。a 和 b 是 m 两边的两个元素，且 b > a。要使 a 和 b 相等，它们总共移动的次数为 b - a，这个值等于 (b - m) + (m - a)，也就是把这两个数移动到中位数的移动次数。
//
// 设数组长度为 N，则可以找到 N/2 对 a 和 b 的组合，使它们都移动到 m 的位置。
func minMoves2(nums []int) int {
	sort.Ints(nums)
	var ret int
	for l, h := 0, len(nums)-1; l < h; l, h = l+1, h-1 {
		ret += nums[h] - nums[l]
	}
	return ret
}

func main() {
	fmt.Println(minMoves2([]int{1, 2, 3}))
}
