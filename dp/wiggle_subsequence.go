package main

import "fmt"

// For every position in the array, there are only three possible statuses for it.
//    up position, it means nums[i] > nums[i-1]
//    down position, it means nums[i] < nums[i-1]
//    equals to position, nums[i] == nums[i-1]
// {{So we can use two arrays up[] and down[] to record the max wiggle sequence length so far at index i.}}
// If nums[i] > nums[i-1], that means it wiggles up. the element before it must be a down position. so up[i] = down[i-1] + 1; down[i] keeps the same with before.
// If nums[i] < nums[i-1], that means it wiggles down. the element before it must be a up position. so down[i] = up[i-1] + 1; up[i] keeps the same with before.
// If nums[i] == nums[i-1], that means it will not change anything becasue it didn’t wiggle at all. so both down[i] and up[i] keep the same.
//
// In fact, we can reduce the space complexity to O(1), but current way is more easy to understanding.
func wiggleMaxLength(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	up, down := make([]int, n), make([]int, n)
	up[0], down[0] = 1, 1
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] > nums[i-1]:
			up[i] = down[i-1] + 1
			down[i] = down[i-1]
		case nums[i] < nums[i-1]:
			down[i] = up[i-1] + 1
			up[i] = up[i-1]
		default:
			up[i] = up[i-1]
			down[i] = down[i-1]
		}
	}
	if up[n-1] > down[n-1] {
		return up[n-1]
	}
	return down[n-1]
}

func wiggleMaxLengthO1(nums []int) int {
	n := len(nums)

	if n == 0 {
		return 0
	}

	up, down := 1, 1
	for i := 1; i < len(nums); i++ {
		switch {
		case nums[i] > nums[i-1]:
			up = down + 1
		case nums[i] < nums[i-1]:
			down = up + 1
		}
	}
	if up > down {
		return up
	}
	return down
}

func main() {
	fmt.Println(wiggleMaxLength([]int{1, 7, 4, 9, 2, 5}))
	fmt.Println(wiggleMaxLengthO1([]int{1, 7, 4, 9, 2, 5}))
}
