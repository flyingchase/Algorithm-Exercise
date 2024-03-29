package qiniuyun

// 判断连通图是否有环的方式
// bfs
// 宽度优先遍历，将遍历过的结点记录染色，再从已经遍历过得结点集合取出遍历所能链接的结点，
// 若已经染色则 false 有环
func isConnected(N int, tu [][]int) bool {
	if len(tu) == 0 {
		return false
	}
	res, isUsed, queue := true, make([]bool, N), make([]int, 0)
	queue = append(queue)
}
