package backtrack

// M-17-电话号码的字母组合

var m = map[int]string{0: "", 1: "", 2: "abc", 3: "def", 4: "ghi", 5: "jkl", 6: "mno", 7: "pqrs", 8: "tuv", 9: "wxyz"}

func letterCombinations(digits string) []string {
	res := []string{}

	if len(digits) == 0 {
		return res
	}
	temp := []byte{}
	dfsLetterCombinations(digits, res, temp, 0)
	return res
}
func dfsLetterCombinations(digits string, res []string, temp []byte, index int) {
	if len(temp) == len(digits) {
		res = append(res, string(temp))
		return
	}

	tempIndex := digits[index] - '0'
	m[tempIndex]
}
