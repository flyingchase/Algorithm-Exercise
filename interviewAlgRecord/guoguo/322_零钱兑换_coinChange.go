package guoguo

import "sort"

// 完全背包
func coinChange(coins []int, amount int) int {
	size := len(coins)
	dp := make([][]int, size+1)
	for i := 0; i < size+1; i++ {
		dp[i] = make([]int, amount+1)
	}
	for i := 0; i < size+1; i++ {
		for j := 0; j < amount+1; j++ {
			// esp
			if i == 0 {
				dp[i][j] = amount + 1
			} else if j == 0 {
				dp[i][j] = 0
			} else if j >= coins[i-1] {
				dp[i][j] = min(dp[i-1][j], dp[i][j-coins[i-1]]+1)
			} else {
				dp[i][j] = dp[i-1][j]
			}
		}
	}
	if dp[size][amount] == amount+1 {
		return -1
	}

	return dp[size][amount]

}

func min(i, j int) int {

	if i > j {
		return j
	}
	return i

}

func bfs(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}

	queue := make([]int, 0)
	queue = append(queue, amount)
	visited := make([]bool, amount+1)

	sort.Ints(coins)

	count := 1

	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]
			queue = queue[1:]

			for _, coin := range coins {
				next := node - coin
				if next == 0 {
					return count
				}
				if next < 0 {
					break
				}

				if visited[next] {
					continue
				}

				queue = append(queue, next)
				visited[next] = true

			}
		}
	}

	return count

}
