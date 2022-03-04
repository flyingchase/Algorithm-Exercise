package interviewprepare

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	dfsGenerateParethesie(n, n, "", &res)
	return res
}
func dfsGenerateParethesie(l, r int, cur string, res *[]string) {
	if l == 0 && r == 0 {
		*res = append(*res, cur)
		return
	}
	if l > 0 {
		dfsGenerateParethesie(l-1, r, cur+"(", res)
	}
	if r > 0 && l < r {
		dfsGenerateParethesie(l, r-1, cur+")", res)
	}
}
