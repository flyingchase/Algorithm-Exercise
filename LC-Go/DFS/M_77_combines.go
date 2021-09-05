package DFS

/*

//给定两个整数 n 和 k，返回范围 [1, n] 中所有可能的 k 个数的组合。
*/

func combine(n int, k int) [][]int {
	if n < 1 || k > n||k<=0 {
		return [][]int {}
	}
	cur, res, index := []int{}, [][]int{}, 1
	dfs(n,cur,&res,index,k)
	return res
}

func dfs(n int, cur []int, res *[][]int, index int, k int) {
	if len(cur) == k {
		tmp := make([] int,len(cur))
		copy(tmp,cur)
		*res=append(*res,tmp)
		return
	}
	// 由于 cur 长度==k 时候被截断 i 最多到 n -(k-len(cur))+1
	for i := index; i <= n; i++ {
		cur=append(cur,i)
		dfs(n,cur,res,i+1,k)
		cur=cur[ : len(cur) - 1]
	}
}
