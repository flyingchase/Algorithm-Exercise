package days

// 最小路径和
func minPathSum(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}
	m, n := len(grid), len(grid[0])
	// dp[i][j]represent minPathSum index i,j
	// dp为 grid 的副本，累加中+=要求
	dp := make([][]int, m)
	for i := 0; i < m; i++ {
		dp[i] = append([]int{}, grid[i]...)
	}
	// board situation
	// 首为 grid网格横移
	for i := 1; i < m; i++ {
		dp[i][0] += dp[i-1][0]
	}
	// 首为 grid 网格竖移
	for j := 1; j < n; j++ {
		dp[0][j] += dp[0][j-1]
	}
	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			// 为左侧和上方最小值累加
			dp[i][j] += min(dp[i-1][j], dp[i][j-1])
		}
	}
	return dp[m-1][n-1]
}

// 零钱兑换
func coinChange(coins []int, amount int) int {
	// 完全背包问题
}

// 第一个缺失的正数
func firstMissingPositive(nums []int) int {
	numMap := make(map[int]int, len(nums))
	// 将 num存在 map 中 以 kv 均为 num 的形式
	for _, num := range nums {
		numMap[num] = num
	}
	for index := 1; index < len(nums)+1; index++ {
		// 正数遍历中找到对应的第一个 v 不存在map 内即为所求
		if _, ok := numMap[index]; !ok {
			return index
		}
	}
	// 缺失的正数最大是数组长度加一
	return len(nums) + 1
}
