package golang

import "sort"

// https://leetcode.com/problems/group-anagrams/discuss/19176/Share-my-short-JAVA-solution
func groupAnagrams(strs []string) [][]string {
	if len(strs) == 0 {
		return nil
	}

	m := make(map[string][]string)
	for _, str := range strs {
		b := []byte(str)
		sort.Slice(b, func(i, j int) bool {
			return b[i] < b[j]
		})
		m[string(b)] = append(m[string(b)], str)
	}

	result := make([][]string, 0, len(m))
	for _, strings := range m {
		result = append(result, strings)
	}
	return result
}
