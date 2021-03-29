package golang

func lengthOfLastWord(s string) int {
	l, r := len(s), len(s)
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] == ' ' {
			if l != r {
				break
			}

			l--
			r--
			continue
		}

		l = i
	}
	return r - l
}
