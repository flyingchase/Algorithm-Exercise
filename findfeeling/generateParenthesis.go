package findfeeling

// 括号生成
func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}

	res := []string{}

	dfsGenerate(n, n, "", &res)
	return res

}

func dfsGenerate(l, r int, curStr string, res *[]string) {
	if l == r && l == 0 {
		*res = append(*res, curStr)
		return
	}
	if l > 0 {
		dfsGenerate(l-1, r, curStr+"(", res)
	}
	if r > 0 && l < r {
		dfsGenerate(l, r-1, curStr+")", res)
	}
}
