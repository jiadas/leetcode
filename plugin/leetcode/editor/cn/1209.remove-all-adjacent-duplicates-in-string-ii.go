package golang

// leetcode submit region begin(Prohibit modification and deletion)
func removeDuplicates(s string, k int) string {
	return doRemoveDuplicates2(s, k)
}

func doRemoveDuplicates1(s string, k int) string {
	if s == "" || len(s) < k {
		return s
	}

	tmp := s
	var l, r int
	for i := 0; i < len(s); i++ {
		if r-l+1 == k {
			tmp = s[:l] + s[r+1:]
			break
		}

		if i < len(s)-1 && s[i] == s[i+1] {
			r++
		} else {
			l, r = i+1, i+1
		}
	}

	if tmp == s {
		return s
	}

	return doRemoveDuplicates1(tmp, k)
}

func doRemoveDuplicates2(s string, k int) string {
	sb := []byte(s)
	count := make([]int, len(s))
	for i := 0; i < len(sb); i++ {
		if i == 0 || sb[i] != sb[i-1] {
			count[i] = 1
		} else {
			count[i] = count[i-1] + 1
			if count[i] == k {
				sb = append(sb[:i-k+1], sb[i+1:]...)
				i = i - k
			}
		}
	}
	return string(sb)
}

// leetcode submit region end(Prohibit modification and deletion)
