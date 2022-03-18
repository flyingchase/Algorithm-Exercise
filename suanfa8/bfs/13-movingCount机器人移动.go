package bfs

func movingCount(m int, n int, k int) int {

	return 0
}

/*
func movingCount(m int, n int, k int) int {
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	return dfs(m, n, 0, 0, k, dp)
}

func dfs(m, n, i, j, k int, dp [][]int) int {
	if i < 0 || j < 0 || i >= m || j >= n || dp[i][j] == 1 || (sumPos(i)+sumPos(j)) > k {
		return 0
	}

	dp[i][j] = 1

	sum := 1
	sum += dfs(m, n, i, j+1, k, dp)
	sum += dfs(m, n, i, j-1, k, dp)
	sum += dfs(m, n, i+1, j, k, dp)
	sum += dfs(m, n, i-1, j, k, dp)
	return sum
}

// 求所有位之和
func sumPos(n int) int {
	var sum int

	for n > 0 {
		sum += n % 10
		n = n / 10
	}

	return sum
}
*/

/*

const maxsize = 105
var (
    dx = [4]int{1,0,-1,0}
    dy = [4]int{0,1,0,-1}
)

type square struct {
    x, y int
}

func digitsSum(i int) int {
    s := 0

    for i > 0 {
        s += i%10
        i /= 10
    }

    return s
}

func bfs(m, n, k int, dp []int) int {
    vis := [maxsize][maxsize]bool{}
    q := []square{}
    q = append(q, square{0,0})
    vis[0][0] = true
    ans := 1
    for len(q) > 0 {
        cur := q[0]
        q = q[1:]
        for i := 0; i < 4; i++ {
            newX, newY := cur.x+dx[i], cur.y+dy[i]
            if (newX>=0 && newX<m && newY>=0 && newY<n) &&
            (!vis[newX][newY] && dp[newX]+dp[newY]<=k) {
                ans++
                q = append(q, square{newX, newY})
                vis[newX][newY] = true
            }
        }
    }
    return ans
}

func movingCount(m int, n int, k int) int {
    dp := [maxsize]int{} // 预处理，指定数字的数位之和
    cursize := int(math.Max(float64(m), float64(n)))
    for i := 0; i < cursize; i++ {
        dp[i] = digitsSum(i)
    }

    return bfs(m, n, k, dp[:])
}
*/
