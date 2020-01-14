package golang

func romanToInt(s string) int {
	m := map[string]int{
		"I":  1,
		"IV": 4,
		"V":  5,
		"IX": 9,
		"X":  10,
		"XL": 40,
		"L":  50,
		"XC": 90,
		"C":  100,
		"CD": 400,
		"D":  500,
		"CM": 900,
		"M":  1000,
	}
	var result int
	for i := len(s) - 1; i >= 0; {
		t := m[string(s[i])]
		if i-1 >= 0 {
			s, ok := m[s[i-1:i+1]]
			if ok {
				t = s
				i--
			}
		}
		result += t
		i--
	}
	return result
}

// https://leetcode.com/problems/roman-to-integer/discuss/6547/Clean-O(n)-c%2B%2B-solution
func romanToIntFD(s string) int {
	m := map[string]int{
		"I": 1,
		"V": 5,
		"X": 10,
		"L": 50,
		"C": 100,
		"D": 500,
		"M": 1000,
	}
	var result int
	if len(s) > 0 {
		result = m[string(s[len(s)-1])]
	}
	if len(s) > 1 {
		for i := len(s) - 2; i >= 0; i-- {
			cur := m[string(s[i])]
			behind := m[string(s[i+1])]
			if cur < behind {
				result -= cur
			} else {
				result += cur
			}
		}
	}
	return result
}
