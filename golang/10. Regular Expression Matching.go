package golang

// https://leetcode.com/problems/regular-expression-matching/discuss/5651/Easy-DP-Java-Solution-with-detailed-Explanation/238767
func isMatch(s string, p string) bool {
	m, n := len(s), len(p)

	// dp[i][j] represents if the 1st i characters in s can match the 1st j characters in p
	dp := make([][]bool, m+1)
	for i := range dp {
		dp[i] = make([]bool, n+1)
	}

	// initialization:
	// 1. dp[0][0] = true, since empty string matches empty pattern
	dp[0][0] = true

	// 2. dp[i][0] = false(which is default value of the boolean array) since empty pattern cannot match non-empty string
	// 3. dp[0][j]: what pattern matches empty string ""? It should be #*#*#*#*...
	// as we can see, the length of pattern should be even && the character at the even position should be *,
	// thus for odd length, dp[0][j] = false which is default. So we can just skip the odd position, i.e. j starts from 2, the interval of j is also 2.
	// and notice that the length of repeat sub-pattern #* is only 2, we can just make use of dp[0][j - 2] rather than scanning j length each time
	// for checking if it matches #*#*#*#*.
	for j := 2; j < n+1; j += 2 {
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}

	// Induction rule is very similar to edit distance, where we also consider from the end. And it is based on what character in the pattern we meet.
	// 1. if p.charAt(j) == s.charAt(i), dp[i][j] = dp[i - 1][j - 1]
	//    ######a(i)
	//    ####a(j)
	// 2. if p.charAt(j) == '.', dp[i][j] = dp[i - 1][j - 1]
	// 	  #######a(i)
	//    ####.(j)
	// 3. if p.charAt(j) == '*':
	//    3.1. if p.charAt(j - 1) != '.' && p.charAt(j - 1) != s.charAt(i), then b* is counted as empty. dp[i][j] = dp[i][j - 2]
	//       #####a(i)
	//       ####b*(j)
	//    3.2. if p.charAt(j - 1) == '.' || p.charAt(j - 1) == s.charAt(i):
	//       ######a(i)
	//       ####.*(j)
	//
	// 	  	 #####a(i)
	//    	 ###a*(j)
	//      3.2.1 if p.charAt(j - 1) is counted as empty, then dp[i][j] = dp[i][j - 2]
	//      3.2.2 if counted as one, then dp[i][j] = dp[i - 1][j - 2]
	//      3.2.3 if counted as multiple, then dp[i][j] = dp[i - 1][j]
	for i := 1; i < m+1; i++ {
		for j := 1; j < n+1; j++ {
			curS, curP := s[i-1], p[j-1]
			if curP == curS || curP == '.' {
				dp[i][j] = dp[i-1][j-1]
			}
			if curP == '*' {
				if j >= 2 {
					preCurP := p[j-2]
					if preCurP != '.' && preCurP != curS {
						dp[i][j] = dp[i][j-2]
					} else {
						dp[i][j] = dp[i][j-2] || dp[i-1][j-2] || dp[i-1][j]
					}
				}
			}
		}
	}

	return dp[m][n]
}
