package golang

import "sort"

// https://leetcode.com/problems/combination-sum/discuss/16502/A-general-approach-to-backtracking-questions-in-Java-(Subsets-Permutations-Combination-Sum-Palindrome-Partitioning)
func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationWithIndex(candidates, nil, 0, target)
}

func combinationWithIndex(candidates, tmpResult []int, index, remain int) [][]int {
	var result [][]int
	switch {
	case remain == 0:
		// copy the tmpResult
		result = append(result, append([]int(nil), tmpResult...))
	case remain > 0:
		for i := index; i < len(candidates); i++ {
			candidate := candidates[i]
			if remain < candidate {
				break
			}

			r := combinationWithIndex(candidates, append(tmpResult, candidate), i, remain-candidate)
			result = append(result, r...)
		}
	}

	return result
}
