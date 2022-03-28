package bfs

// 选修课程
// 判断有向图是否为有向无环图,拓扑排序
func canFinish(numCourses int, prerequisites [][]int) bool {
	// 有向图
	T := make([][]int, numCourses)
	// 每个节点的入度
	in_degree := make([]int, numCourses)
	res := make([]int, 0, numCourses)
	for _, v := range prerequisites {
		// v[1]-->v[0]
		T[v[1]] = append(T[v[1]], v[0])
		in_degree[v[0]]++
	}
	queue := []int{}
	// 度==0 入队
	for i := 0; i < numCourses; i++ {
		if in_degree[i] == 0 {
			queue = append(queue, i)
		}
	}
	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]
		res = append(res, node)
		for _, v := range T[node] {
			in_degree[v]--
			// 相邻结点的入度=0 则入队
			if in_degree[v] == 0 {
				queue = append(queue, v)
			}
		}
	}
	return len(res) == numCourses
}
