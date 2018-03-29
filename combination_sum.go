package main

import (
	"fmt"
	"sort"
)

func combinationSum(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	return doCombinationSum(candidates, nil, 0, target, nil)
}

func doCombinationSum(candidates []int, c []int, sum int, target int, result [][]int) [][]int {
	if sum == target {
		return append(result, append([]int(nil), c...))
	}
	if sum > target {
		return result
	}
	for i := 0; i < len(candidates); i++ {
		if sum+candidates[i] <= target {
			r := append([]int(nil), candidates[i:]...)
			result = append(result, doCombinationSum(r, append(c, candidates[i]), sum+candidates[i], target, nil)...)
		} else {
			break
		}

	}
	return result
}

func main() {
	fmt.Println(combinationSum([]int{2, 3, 6, 7}, 7))
}
