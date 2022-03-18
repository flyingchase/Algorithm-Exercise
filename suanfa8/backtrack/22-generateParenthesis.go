package backtrack

// 括号生成
func generateParenthesis(n int) []string {
	if n == 0 {
		return nil
	}
	res := []string{}
	backtrackGenerateparenthesis(n, n, "", &res)
	return res
}
func backtrackGenerateparenthesis(l, r int, cur string, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, cur)
		return
	}
	if l > 0 {
		backtrackGenerateparenthesis(l-1, r, cur+"(", res)
	}
	if r > 0 && l < r {
		backtrackGenerateparenthesis(l, r-1, cur+")", res)
	}
}
