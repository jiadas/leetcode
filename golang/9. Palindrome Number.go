package golang

// https://leetcode.com/problems/palindrome-number/discuss/5127/9-line-accepted-Java-code-without-the-need-of-handling-overflow
// compare half of the digits in x, so don't need to deal with overflow.
// 与 isPalindromeWithoutConcernOverflow 相比快了 20ms
func isPalindrome(x int) bool {
	if x < 0 || (x != 0 && x%10 == 0) {
		return false
	}
	var rev int
	for x > rev {
		rev = rev*10 + x%10
		x = x / 10
	}
	return x == rev || x == rev/10
}

// https://leetcode.com/problems/palindrome-number/discuss/5127/9-line-accepted-Java-code-without-the-need-of-handling-overflow/5915
// When the reversed number overflows, it will becomes negative number which will return false when compared with x.
func isPalindromeWithoutConcernOverflow(x int) bool {
	if x < 0 {
		return false
	}
	if x == 0 {
		return true
	}
	if x%10 == 0 {
		return false
	}
	y := x
	var rev int
	for y > 0 {
		rev = rev*10 + y%10
		y /= 10
	}
	return x == rev
}
