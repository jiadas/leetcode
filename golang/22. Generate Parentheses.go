package golang

// https://leetcode.com/problems/generate-parentheses/discuss/10100/Easy-to-understand-Java-backtracking-solution
func generateParenthesis(n int) []string {
	return backtrackGenerateParenthesis(nil, "", 0, 0, n)
}

func backtrackGenerateParenthesis(result []string, str string, open, close, max int) []string {
	if len(str) == 2*max {
		return append(result, str)
	}
	if open < max {
		result = backtrackGenerateParenthesis(result, str+"(", open+1, close, max)
	}
	if close < open {
		result = backtrackGenerateParenthesis(result, str+")", open, close+1, max)
	}
	return result
}
