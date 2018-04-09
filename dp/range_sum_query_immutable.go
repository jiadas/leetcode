package main

import "fmt"

type NumArray struct {
	nums []int
	sum  []int // sum[i] 表示 [0,i] 子集的元素和
}

func Constructor(nums []int) NumArray {
	sum := make([]int, len(nums))
	for i := 0; i < len(nums); i++ {
		if i == 0 {
			sum[i] = nums[i]
		} else {
			sum[i] = sum[i-1] + nums[i]
		}
	}
	return NumArray{nums: nums, sum: sum}
}

func (this *NumArray) SumRange(i int, j int) int {
	if i == 0 {
		return this.sum[j]
	}
	return this.sum[j] - this.sum[i-1]
}

/**
 * Your NumArray object will be instantiated and called as such:
 * obj := Constructor(nums);
 * param_1 := obj.SumRange(i,j);
 */

func main() {
	na := Constructor([]int{-2, 0, 3, -5, 2, -1})
	fmt.Println(na.SumRange(0, 2))
	fmt.Println(na.SumRange(2, 5))
	fmt.Println(na.SumRange(0, 5))
}
