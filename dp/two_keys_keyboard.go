package main

import "fmt"

// Now, check the solution.
// Eg: n=6
// To get 6, we need to copy 3 'A’s two time. (2)
// To get 3 'A’s, copy the 1 ‘A’ three times. (3)
// So the answer for 6 is 5
//
// Now, take n=9.
// We need the lowest number just before 9 such that (9 % number =0). So the lowest number is 3.
// So 9%3=0. We need to copy 3 'A’s three times to get 9. (3)
// For getting 3 'A’s, we need to copy 1 ‘A’ three times. (3)
// So the answer is 6
//
// Finally to analyse the below code, take n=81.
// To get 81 we check
// if (81 % 2 ==0) No
// if (81 % 3 ==0) Yes
// So we need to copy 81/3 = 27 'A’s three times (3)
// Now to get 27 'A’s, we need to copy 27/3= 9 'A’s three times (3)
// To get 9 'A’s, we need to copy 9/3=3 'A’s three times (3)
// And to get 3 'A’s, we need to copy 3/3=1 'A’s three times (3)
// Final answer is 3+3+3+3 = 12
//
// Last Example, n=18
// 18/2 = 9 Copy 9 'A’s 2 times (2)
// 9/3=3 Copy 3 'A’s 3 times (3)
// 3/3=1 Copy 1’A’s 3 times (3)
// Answer: 2+3+3= 8
func minSteps(n int) int {
	var ret int
	for i := 2; i <= n; i++ {
		for n%i == 0 {
			ret += i
			n = n / i
		}
	}
	return ret
}

func minStepsDP(n int) int {
	if n == 0 || n == 1 {
		return 0
	}
	dp := make([]int, n+1) // dp[i] 表示获得 i 个 'A' 需要的最少步数
	for i := 2; i <= n; i++ {
		dp[i] = i
		for j := i - 1; j > 1; j-- {
			if i%j == 0 {
				dp[i] = dp[j] + i/j
				break
			}
		}
	}
	return dp[n]
}

func main() {
	fmt.Println(minSteps(3))
	fmt.Println(minStepsDP(81))
}
