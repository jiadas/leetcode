package golang

// https://leetcode.com/problems/find-k-closest-elements/discuss/106426/JavaC%2B%2BPython-Binary-Search-O(log(N-K)-%2B-K)
// 官方题解：
// https://leetcode-cn.com/problems/find-k-closest-elements/solution/zhao-dao-kge-zui-jie-jin-de-yuan-su-by-leetcode/
// 官方题解简单易懂，只是多了为缩小 k 范围所消耗的时间
func findClosestElements(arr []int, k int, x int) []int {
	left, right := 0, len(arr)-k
	var mid int
	for left < right {
		mid = left + (right-left)>>1
		if x-arr[mid] > arr[mid+k]-x {
			left = mid + 1
		} else {
			right = mid
		}
	}
	return arr[left : left+k]
}
