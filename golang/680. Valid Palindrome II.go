package golang

// 想出了用双指针来判断回文，也知道当 low 和 high 不相同是必须要删一个，但是在具体处理删除时却拿捏不准了
// 官方题解的处理是：当这两个子串中至少有一个是回文串时，就说明原始字符串删除一个字符之后就以成为回文串。就是删 low 和删 high 都处理，有一个是回文就妥了。
func validPalindrome(s string) bool {
	low, high := 0, len(s)-1
	for low < high {
		if s[low] == s[high] {
			low++
			high--
		} else {
			flag1, flag2 := true, true
			for i, j := low+1, high; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag1 = false
					break
				}
			}
			for i, j := low, high-1; i < j; i, j = i+1, j-1 {
				if s[i] != s[j] {
					flag2 = false
					break
				}
			}
			return flag1 || flag2
		}
	}
	return true
}
