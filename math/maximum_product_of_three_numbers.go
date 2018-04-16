package main

import (
	"fmt"
	"math"
)

// Simply find out the three largest numbers and the two smallest numbers using one pass.
func maximumProduct(nums []int) int {
	max1, max2, max3 := math.MinInt32, math.MinInt32, math.MinInt32
	min1, min2 := math.MaxInt32, math.MaxInt32
	for _, v := range nums {
		if v > max1 {
			max3 = max2
			max2 = max1
			max1 = v
		} else if v > max2 {
			max3 = max2
			max2 = v
		} else if v > max3 {
			max3 = v
		}

		if v < min1 {
			min2 = min1
			min1 = v
		} else if v < min2 {
			min2 = v
		}
	}
	max := max1 * max2 * max3
	if max < min1*min2*max1 {
		max = min1 * min2 * max1
	}
	return max
}

func main() {
	fmt.Println(maximumProduct([]int{1, 2, 3, 4}))
	fmt.Println(maximumProduct([]int{-4, -3, -2, -1, 60}))
}
