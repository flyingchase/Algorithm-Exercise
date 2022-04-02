package days

import "sort"

// 最长有效括号
func longestValidParentheses(s string) int {
	 count := len(s)
    if count == 0 {
        return 0
    } 
    dp := make([]int, count)
    dp[0] = 0
    for i:=1;i<count;i++ {
        if s[i] == ')' {
            if s[i-1] == '(' {
                if i>1 {
                    dp[i] = dp[i-2] + 2
                } else {
                    dp[i] = 2
                }
            } else {
                if i > dp[i-1] {
                    if s[i-dp[i-1]-1] == '(' {
                        if i>=dp[i-1]+2 {
                            dp[i] = dp[i-dp[i-1]-2] + dp[i-1] + 2
                        } else {
                            dp[i] = dp[i-1] + 2 
                        }
                    } 
                }
            }
        } else {
            dp[i] = 0
        } 
    }
    max := 0 
    for _, v := range dp {
        if max < v {
            max = v
        }
    }
    return max
}
// 比较版本号
func compareVersion(version1 string, version2 string) int {
	v1 := strings.Split(version1, ".")
	v2 := strings.Split(version2, ".")

	for len(v1) < len(v2) {
		v1 = append(v1, "0")
	}
	for len(v2) < len(v1) {
		v2 = append(v2, "0")
	}

	l := len(v1)
	for i := 0; i < l; i++ {
		vs1 := strings.TrimLeft(v1[i], "0")
		vs2 := strings.TrimLeft(v2[i], "0")

		for len(vs1) < len(vs2) {
			vs1 = "0" + vs1
		}
		for len(vs2) < len(vs1) {
			vs2 = "0" + vs2
		}

		vl := len(vs1)
		for j := 0; j < vl; j++ {
			if vs1[j] == vs2[j] {
				continue
			}

			if vs1[j] < vs2[j] {
				return -1
			}

			return 1
		}
	}

	return 0
}func combinationSum(candidates []int, target int) [][]int {
	sort.Ints(candidates)
	var ret [][]int
	helper(candidates, target, 0, []int{}, &ret)
	return ret
}

func helper(candidates []int, target, start int, temp []int, ret *[][]int) {
	if target < 0 {
		return
	}
	if target == 0 {
		*ret = append(*ret, append([]int{}, temp...))
		return
	}
	for i := start; i < len(candidates); i++ {
		temp = append(temp, candidates[i])
		helper(candidates, target-candidates[i], i, temp, ret)
		temp = temp[:len(temp)-1]
	}
}
