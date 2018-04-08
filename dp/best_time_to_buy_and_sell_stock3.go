package main

import (
	"fmt"
	"math"
)

// The thinking is simple and is inspired by the best solution from Single Number II (I read through the discussion after I use DP).
// Assume we only have 0 money at first;
// 4 Variables to maintain some interested ‘ceilings’ so far:
// The maximum of if we’ve just buy 1st stock, if we’ve just sold 1nd stock, if we’ve just buy 2nd stock, if we’ve just sold 2nd stock.
// Very simple code too and work well. I have to say the logic is simple than those in Single Number II.
func maxProfit3(prices []int) int {
	n := len(prices)
	if n <= 1 {
		return 0
	}

	hold1, hold2 := math.MinInt32, math.MinInt32
	release1, release2 := 0, 0
	// Assume we only have 0 money at first
	for _, price := range prices {
		release2 = isMax(release2, hold2+price) // The maximum if we've just sold 2nd stock so far.
		hold2 = isMax(hold2, release1-price)    // The maximum if we've just buy  2nd stock so far.
		release1 = isMax(release1, hold1+price) // The maximum if we've just sold 1nd stock so far.
		hold1 = isMax(hold1, -price)            // The maximum if we've just buy  1st stock so far.
	}

	return release2
}

func isMax(a, b int) int {
	if a >= b {
		return a
	}
	return b
}

func main() {
	fmt.Println(maxProfit3([]int{3, 3, 5, 0, 0, 3, 1, 4}))
}
