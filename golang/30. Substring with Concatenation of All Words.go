package golang

// https://leetcode.wang/leetCode-30-Substring-with-Concatenation-of-All-Words.html
func findSubstring(s string, words []string) []int {
	var indexes []int
	wordNum := len(words)
	if wordNum == 0 {
		return indexes
	}
	counts := make(map[string]int, wordNum)
	for _, word := range words {
		counts[word] += 1
	}
	n, wordLen := len(s), len(words[0])
	for i := 0; i < n-wordNum*wordLen+1; i++ {
		seen := make(map[string]int)
		var j int
		for j < wordNum {
			word := s[i+j*wordLen : i+(j+1)*wordLen]
			if c, ok := counts[word]; !ok {
				break
			} else {
				seen[word] += 1
				if seen[word] > c {
					break
				}
			}
			j++
		}
		if j == wordNum {
			indexes = append(indexes, i)
		}
	}
	return indexes
}
