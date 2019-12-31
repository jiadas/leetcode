package main

// 本题主要问题是判断两个字符串是否含相同字符，由于字符串只含有小写字符，总共 26 位，
// 因此可以用一个 32 位的整数来存储每个字符是否出现过。
func maxProduct(words []string) int {
	mark := make([]int, len(words))
	for i, w := range words {
		for j := 0; j < len(w); j++ {
			mark[i] |= 1 << (w[j] - 'a')
		}
	}

	var ret int
	for i := 0; i < len(words); i++ {
		for j := i + 1; j < len(words); j++ {
			if mark[i]&mark[j] == 0 {
				if ret < len(words[i])*len(words[j]) {
					ret = len(words[i]) * len(words[j])
				}
			}
		}
	}

	return ret
}
