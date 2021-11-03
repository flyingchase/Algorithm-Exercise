package firstround

func generateParenthesis(n int) []string {
	if n == 0 {
		return []string{}
	}
	res := []string{}
	dfsGenerateparenthesis(n, n, "", &res)
	return res
}
func dfsGenerateparenthesis(left, right int, curString string, res *[]string) {
	if left == 0 && right == 0 {
		*res = append(*res, curString)
		return
	}
	if left > 0 {
		dfsGenerateparenthesis(left-1, right, curString+"(", res)
	}
	if right > 0 && left < right {
		dfsGenerateparenthesis(left, right+1, curString+")", res)
	}
}
