package main

import "fmt"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	// For nums[0…n-1], 0 and n-1 are neighboring to each other.
	// Basically, there are only three possible cases:
	// (1) rob 0, but leave n-1 untouched  =>  [0,...,n-2] 包含了不抢0的可能
	// (2) leave 0 untouched, rob n-1      =>  [1,...,n-1] 包含了不抢n-1的可能
	// (3) leave both 0 and n-1 untouched. Obviously, case (3) can be covered by case(1) or case (2) in the simple House Robber problem
	return maxRob(doRob(nums, 0, len(nums)-2), doRob(nums, 1, len(nums)-1))
}

func doRob(nums []int, low, high int) int {
	// 用来记录对 当前这户人家 抢 和 不抢 的最大值
	var include, exclude int
	for j := low; j <= high; j++ {
		// 保存对上一户人家 抢 和 不抢 的最大值
		i, e := include, exclude

		// 计算对 当前这户人家 抢 和 不抢 的最大值
		include = exclude + nums[j]
		exclude = maxRob(i, e)
	}
	return maxRob(include, exclude)
}

func maxRob(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(rob([]int{1, 2, 3, 4, 5}))
}
