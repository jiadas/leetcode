package w186

func maxScore(s string) int {
	var maxScore int
	for i := 1; i < len(s); i++ {
		l := count(s[:i], byte('0'))
		r := count(s[i:], byte('1'))
		if l+r > maxScore {
			maxScore = l + r
		}
	}
	return maxScore
}

func count(s string, x byte) int {
	var count int
	for i := 0; i < len(s); i++ {
		if s[i] == x {
			count++
		}
	}
	return count
}
