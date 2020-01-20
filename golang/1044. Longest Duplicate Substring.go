package golang

import (
	"math"
	"strings"
)

// 官方题解：
// https://leetcode-cn.com/problems/longest-duplicate-substring/solution/zui-chang-zhong-fu-zi-chuan-by-leetcode/
// 借鉴了 search 方法，即直接放回可能的最长子串，而不是像官方题解里返回的是下标：
// https://leetcode.com/problems/longest-duplicate-substring/discuss/291048/C%2B%2B-solution-using-Rabin-Karp-and-binary-search-with-detailed-explaination
// Discuss 里用 Go 完美 AC 的提交：
// https://leetcode-cn.com/problems/longest-duplicate-substring/comments/93036
func longestDupSubstring(S string) string {
	var result string
	left, right := 1, len(S)
	for left != right {
		L := left + (right-left)/2
		tmp := search(L, S)
		// tmp := RabinKarp(L, S)
		if len(tmp) != 0 {
			if len(tmp) > len(result) {
				result = tmp
			}
			left = L + 1
		} else {
			right = L
		}
	}
	tmp := search(left-1, S)
	if len(tmp) > len(result) {
		result = tmp
	}
	return result
}

// Rabin-Karp with polynomial rolling hash.
// Search a substring of given length
// that occurs at least 2 times.
func search(L int, S string) string {
	a, modulus := 26, math.MaxInt32
	// compute the hash of string S[:L]
	var h int
	for i := 0; i < L; i++ {
		h = (h*a + int(S[i]-'a')) % modulus
	}
	// already seen hashes of strings of length L
	seen := map[int]struct{}{h: {}}
	// const value to be used often : a**L % modulus
	aL := 1
	for i := 0; i < L; i++ {
		aL = (aL * a) % modulus
	}

	for i := 1; i < len(S)-L+1; i++ {
		// compute rolling hash in O(1) time
		h = (h*a - aL*int(S[i-1]-'a')%modulus + modulus) % modulus
		h = (h + int(S[i+L-1]-'a')) % modulus
		// 用下面这种方式计算 h 不能 pass
		// h = (h*a - int(S[i-1]-'a')*aL + int(S[i+L-1]-'a')) % modulus
		// 如果不用 strings.Contains 检查下 S[i:i+L] 是否包含在 S[:i+L-1] 里也不能 pass，
		// 一定要是 i+L-1, 意思是除了 S[i:i+L] 外是否还有跟 S[i:i+L] 相同的子串！！！
		// 因为存在 hash 值相同，但是字符串不同的情况
		// if _, ok := seen[h]; ok {
		if _, ok := seen[h]; ok && strings.Contains(S[:i+L-1], S[i:i+L]) {
			return S[i : i+L]
		}
		seen[h] = struct{}{}
	}
	return ""
}

func RabinKarp(length int, s string) string {
	m := make(map[int]bool)
	r, mod := 256, 6*(1<<20)+1
	now := 0
	for i := 0; i < length; i++ {
		now = ((now * r) + int(s[i])) % mod
	}
	rm := 1
	for i := 1; i < length; i++ {
		rm = (rm * r) % mod
	}
	m[now] = true
	for i := length; i < len(s); i++ {
		now = (now - rm*int(s[i-length])%mod + mod) % mod
		now = (now*r + int(s[i])) % mod
		if m[now] && strings.Contains(s[:i], s[i-length+1:i+1]) {
			return s[i-length+1 : i+1]
		}
		m[now] = true
	}
	return ""
}
