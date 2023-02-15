package golang

// leetcode submit region begin(Prohibit modification and deletion)
// 「滚动数组思想」是一种常见的动态规划优化方法，在我们的题目中已经多次使用到，
// 例如「剑指 Offer 46. 把数字翻译成字符串」、「70. 爬楼梯」等，
// 当我们定义的状态在动态规划的转移方程中只和某几个状态相关的时候，就可以考虑这种优化方法，目的是给空间复杂度「降维」。
// 如果你还不知道什么是「滚动数组思想」，一定要查阅相关资料进行学习哦。
// 👆官方题解的代码中给出了使用「滚动数组思想」优化后的实现。对「滚动数组思想」还不了解，暂时理解不了。先用空间复杂度O(mn)来解
func uniquePathsWithObstacles(obstacleGrid [][]int) int {
	// [0][0]直接就是障碍物我是万万想不到的[震惊]
	if obstacleGrid[0][0] == 1 {
		return 0
	}

	m, n := len(obstacleGrid), len(obstacleGrid[0])
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = make([]int, n)
		//if obstacleGrid[i][0] == 0 {
		//	dp[i][0] = 1
		//}
	}

	// 👇🏻这样的初始化是不对的，如果第一行里有障碍物，那么障碍物之后的点dp都应该赋值为 0
	//for i := 0; i < n; i++ {
	//	if obstacleGrid[0][i] == 0 {
	//		dp[0][i] = 1
	//	}
	//}
	for i := 0; i < m && obstacleGrid[i][0] == 0; i++ {
		dp[i][0] = 1
	}
	for i := 0; i < n && obstacleGrid[0][i] == 0; i++ {
		dp[0][i] = 1
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			if obstacleGrid[i][j] == 0 {
				dp[i][j] = dp[i-1][j] + dp[i][j-1]
			}
		}
	}

	return dp[m-1][n-1]
}

//leetcode submit region end(Prohibit modification and deletion)
