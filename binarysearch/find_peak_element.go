package main

import "fmt"

// https://leetcode.com/problems/find-peak-element/discuss/50239/Java-solution-and-explanation-using-invariants
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	l, h := 0, len(nums)-1
	for l < h {
		mid := l + (h-l)>>1
		if nums[mid] < nums[mid+1] {
			l = mid + 1
		} else {
			h = mid
		}
	}
	return l
}

func main() {
	fmt.Println(findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}))
}
