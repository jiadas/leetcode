package golang

// https://leetcode.com/problems/letter-combinations-of-a-phone-number/discuss/8064/My-java-solution-with-FIFO-queue
func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return nil
	}
	mapping := []string{"0", "1", "abc", "def", "ghi", "jkl", "mno", "pqrs", "tuv", "wxyz"}
	result := []string{""}
	for i, digit := range digits {
		x := digit - '0'
		for len(result[0]) == i {
			var t string
			t, result = result[0], result[1:]
			for _, s := range mapping[x] {
				result = append(result, t+string(s))
			}
		}
	}
	return result
}
