package main

import (
	"fmt"
	"sort"
)

func combinationSum2(candidates []int, target int) [][]int {
	if len(candidates) == 0 {
		return nil
	}
	sort.Ints(candidates)
	return doCombinationSum2(candidates, nil, 0, target, nil)
}

func doCombinationSum2(candidates []int, c []int, sum int, target int, result [][]int) [][]int {
	if sum == target {
		return append(result, append([]int(nil), c...))
	}
	if sum > target {
		return result
	}
	for i := 0; i < len(candidates); i++ {
		if i > 0 && candidates[i] == candidates[i-1] {
			continue
		}

		if sum+candidates[i] <= target {
			r := append([]int(nil), candidates[i+1:]...)
			result = append(result, doCombinationSum2(r, append(c, candidates[i]), sum+candidates[i], target, nil)...)
		} else {
			break
		}

	}
	return result
}

func main() {
	// 1,1,2,5,6,7,8,10
	fmt.Println(combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
}
