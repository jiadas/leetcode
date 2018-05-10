package main

// 利用 x ^ 1s = ~x 的特点，可以将位级表示翻转；
// 利用 x ^ x = 0 的特点，可以将三个数中重复的两个数去除，只留下另一个数。
func missingNumber(nums []int) int {
	var ret int
	// 相当于让 nums 和 0...n 这两组数字进行进行异或，去除两组数中重复的，留下的那个数就是缺失的
	for i := 0; i < len(nums); i++ {
		ret ^= i ^ nums[i]
	}
	// 在遍历 nums 时，为了不越界，i 最大只能到 n-1
	// 最好还差和 n 进行异或
	return ret ^ len(nums)
}
