package golang

import "sort"

func combinationSum2(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	return combinationWithIndex2(candidates, nil, 0, target)
}

func combinationWithIndex2(candidates, tmpResult []int, index, remain int) [][]int {
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

			r := combinationWithIndex2(candidates, append(tmpResult, candidate), i+1, remain-candidate)
			result = append(result, r...)

			for i+1 < len(candidates) && candidates[i+1] == candidate {
				i++
			}
		}
	}

	return result
}
