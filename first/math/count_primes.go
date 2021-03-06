package main

import "fmt"

// https://zh.wikipedia.org/wiki/%E5%9F%83%E6%8B%89%E6%89%98%E6%96%AF%E7%89%B9%E5%B0%BC%E7%AD%9B%E6%B3%95
// 先用2去筛，即把2留下，把2的倍数剔除掉；再用下一个素数，也就是3筛，把3留下，把3的倍数剔除掉；接下去用下一个素数5筛，把5留下，把5的倍数剔除掉；不断重复下去......。
//
// 步骤
// 埃拉托斯特尼筛法
// 详细列出算法如下：
//
// 列出2以后的所有序列：
// 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
// 标出序列中的第一个质数，也就是2，序列变成：
// 2 3 4 5 6 7 8 9 10 11 12 13 14 15 16 17 18 19 20 21 22 23 24 25
// 将剩下序列中，划摽2的倍数（用红色标出），序列变成：
// 2 3 [4] 5 [6] 7 [8] 9 [10] 11 [12] 13 [14] 15 [16] 17 [18] 19 [20] 21 [22] 23 [24] 25
// 如果现在这个序列中最大数小于最后一个标出的素数的平方，那么剩下的序列中所有的数都是质数，否则回到第二步。
// 本例中，因为25大于2的平方，我们返回第二步：
// 剩下的序列中第一个质数是3，将主序列中3的倍数划出（红色），主序列变成：
// 2 3 [4] 5 [6] 7 [8] [9] [10] 11 [12] 13 [14] [15] [16] 17 [18] 19 [20] [21] [22] 23 [24] 25
// 我们得到的质数有：2，3
// 25仍然大于3的平方，所以我们还要返回第二步：
// 现在序列中第一个质数是5，同样将序列中5的倍数划出，主序列成了：
// 2 3 [4] 5 [6] 7 [8] [9] [10] 11 [12] 13 [14] [15] [16] 17 [18] 19 [20] [21] [22] 23 [24] [25]
// 我们得到的质数有：2, 3, 5 。
// 因为25等于5的平方，结束循环
// 结论：去掉红色的数字，2到25之间的质数是：2, 3, 5, 7, 11, 13, 17, 19, 23。
func countPrimes(n int) int {
	notPrime := make([]bool, n)
	var count int
	for i := 2; i < n; i++ {
		if notPrime[i] == false {
			count++
			// 从 i * i 开始，因为如果 k < i，那么 k * i 在之前就已经被去除过了
			// 虽然
			// for j := 1; j*i < n; j++ {
			//    notPrime[j*i] = true
			// }
			// 更直观一点
			for j := i * i; j < n; j += i {
				notPrime[j] = true
			}
		}
	}
	return count
}

func main() {
	fmt.Println(countPrimes(25))
	fmt.Println(countPrimes(23))
}
