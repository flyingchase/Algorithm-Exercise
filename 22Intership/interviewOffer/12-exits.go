package interviewoffer

// 二维数组中查询是否可以组成 word
func exist(board [][]byte, word string) bool {
	if len(board) == 0 || len(board[0]) == 0 {
		return false
	}
	m, n := len(board), len(board[0])
	words := []byte(word)
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if dfs(board, words, i, j, 0) {
				return true
			}
		}
	}
	return false
}
func dfs(board [][]byte, words []byte, i, j int, count int) bool {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[0]) {
		return false
	}
	if board[i][j] == words[count] {
		if count == len(words)-1 {
			return true
		}
		tmp := board[i][j]
		// 将 board 置为已经访问过
		board[i][j] = '#'
		found := dfs(board, words, i+1, j, count+1) || dfs(board, words, i-1, j, count+1) || dfs(board, words, i, j+1, count+1) || dfs(board, words, i, j-1, count+1)
		board[i][j] = tmp
		return found
	}
	return false
}
