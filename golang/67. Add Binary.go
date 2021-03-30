package golang

// 题解：https://leetcode-cn.com/problems/add-binary/solution/er-jin-zhi-qiu-he-by-leetcode-solution/
func addBinary(a string, b string) string {
	if len(a) > len(b) {
		a, b = b, a
	}

	ans := make([]byte, len(b)+1)
	c := byte('0')
	i := 0
	for ; i < len(a); i++ {
		ans[len(ans)-1-i], c = doAddBinary(a[len(a)-1-i], b[len(b)-1-i], c)
	}

	for j := len(b) - 1 - i; j >= 0; j-- {
		ans[j+1], c = addTwoBit(b[j], c)
	}

	if c == '1' {
		ans[0] = c
		return string(ans)
	}

	return string(ans[1:])
}

func doAddBinary(a, b, t byte) (r, c byte) {
	r, c = addTwoBit(a, b)
	if c == '1' {
		r, _ = addTwoBit(r, t)
		return r, c
	}
	return addTwoBit(r, t)
}

func addTwoBit(a, b byte) (r, c byte) {
	switch {
	case a == '1' && b == '1':
		return '0', '1'
	case a == '0' && b == '0':
		return '0', '0'
	default:
		return '1', '0'
	}
}
