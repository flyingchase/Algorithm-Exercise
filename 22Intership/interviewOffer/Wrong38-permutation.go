package interviewoffer

func Permutation(s string) []string {
	if len(s) == 0 {
		return nil
	}
	if len(s) == 1 {
		return []string{s}
	}
	res := [][]byte{}
	sSlice := []byte{}
	for i := 0; i < len(s); i++ {
		sSlice = append(sSlice, s[i])
	}
	backtrackPermutation(sSlice, &res, []byte{})
	tmp := []string{}
	for i := 0; i < len(res); i++ {
		tmp[i] = string(res[i])
	}
	return tmp
}
func backtrackPermutation(s []byte, res *[][]byte, cur []byte) {
	if len(cur) == len(s) {
		*res = append(*res, append([]byte{}, cur...))
	}
	for i := 0; i < len(s); i++ {
		if i > 0 && s[i] == s[i-1] {
			continue
		}
		cur = append(cur, s[i])
		backtrackPermutation(s, res, cur)
		cur = cur[:len(cur)-1]
	}
}
