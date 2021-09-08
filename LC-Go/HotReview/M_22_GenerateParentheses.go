package HotReview

/*
有效括号数目 n代表生成括号的对数 返回有效的括号组合（内嵌括号）

dfsGenerate backtrack 会使得括号自动配对
*/

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	res := []string{}
	dfsGenerate(n, n, "", &res)
	return res

}

func dfsGenerate(LeftIndex, RightIndex int, curStr string, res *[]string) {

	if LeftIndex == 0 && RightIndex == 0 {
		*res = append(*res, curStr)
		return
	}

	if LeftIndex > 0 {
		dfsGenerate(LeftIndex-1, RightIndex, curStr+"(", res)
	}
	if RightIndex > 0 && LeftIndex < RightIndex {
		dfsGenerate(LeftIndex, RightIndex-1, curStr+")", res)
	}
}
