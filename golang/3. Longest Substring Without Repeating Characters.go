package golang

func lengthOfLongestSubstring(s string) int {
	if s == "" {
		return 0
	}

	var result int
	for i := range s {
		m := make(map[byte]struct{}, len(s))
		var counter int
		for j := i; j < len(s); j++ {
			c := s[j]
			if _, ok := m[c]; !ok {
				counter++
				m[c] = struct{}{}
				if counter > result {
					result = counter
				}
			} else {
				break
			}
		}
	}
	return result
}

// https://leetcode.com/problems/longest-substring-without-repeating-characters/discuss/1729/11-line-simple-Java-solution-O(n)-with-explanation
func lengthOfLongestSubstringFD(s string) int {
	if s == "" {
		return 0
	}

	m := make(map[byte]int, len(s))
	var max int
	for i, j := 0, 0; i < len(s); i++ {
		if index, ok := m[s[i]]; ok {
			// Because (j) here represents the start index of our string.
			// (i) is the pointer that traverses the whole string.
			// Say we find s.charAt(i) is already in the map,
			// 	case 1: its index < j, then our start index should still be j because it is no longer included in the calculation of the length of the string;
			// 	case 2: index >= j, then our start index should be (index + 1).
			j = maxInt(j, index+1)
		}
		m[s[i]] = i
		max = maxInt(max, i-j+1)
	}
	return max
}
