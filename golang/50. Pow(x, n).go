package golang

// https://leetcode.com/problems/powx-n/discuss/19546/Short-and-easy-to-understand-solution
func myPow(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}
	if n%2 == 0 {
		return myPow(x*x, n/2)
	}
	return x * myPow(x*x, n/2)
}

// 官方题解方法 3：快速幂算法（循环）
// https://leetcode-cn.com/problems/powx-n/solution/powx-n-by-leetcode/
func myPowIter(x float64, n int) float64 {
	if n == 0 {
		return 1
	}
	if n < 0 {
		n = -n
		x = 1 / x
	}

	// x = a + b + c ==> pow(x,n) = pow(x,a) * pow(x,b) * pow(x,c)
	// x = 7 二进制转十进制 7 = 4 + 2 +1
	result := 1.0
	cur := x // pow(x,1)
	for i := n; i > 0; i = i >> 1 {
		if i%2 == 1 { // 等价于判断最后一位是否为 1
			result *= cur
		}
		cur *= cur // pow(x,i)
	}

	return result
}
