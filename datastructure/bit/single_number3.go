package main

// n&(-n) 得到 n 的位级表示中最低的那一位。
// -n 得到 n 的反码加 1，对于二进制表示 10110 100 ，-n 得到 01001100，相与得到 00000100
func singleNumber3(nums []int) []int {
	var diff int
	for _, n := range nums {
		diff ^= n
	}
	// 最右边的 1
	diff &= -diff

	ret := make([]int, 2)
	// 对 nums 按跟 diff 最右边的位置是否为 1 进行分组
	for _, n := range nums {
		if n&diff == 0 { // 判断条件要用 0
			ret[0] ^= n
		} else {
			ret[1] ^= n
		}
	}
	return ret
}
