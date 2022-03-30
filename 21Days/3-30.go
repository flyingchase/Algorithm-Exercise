package days

import (
	"math"
	"sort"
)

// 字符串相乘
func multiply(num1 string, num2 string) string {
	m, n := len(num1), len(num2)
	res := make([]byte, 0, m+n+2)
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j >= 0; j-- {
			tmp := (num1[i]-'0')*(num2[j]-'0') + res[i+j+1]
			res[i+j], res[i+j+1] = tmp/10+res[i+j], tmp%10
		}
	}
	for i := 0; i < len(res); i++ {
		res[i] += '0'
	}
	for i := 0; i < len(res); i++ {
		if res[i] != '0' {
			return string(res[i:])
		}
	}
	return "0"
}

// 子集 不含重复元素
func subsets(nums []int) [][]int {
	if len(nums) == 0 {
		return [][]int{{}}
	}
	sort.Ints(nums)
	res := [][]int{}
	dfsSubsets(&res, []int{}, nums, 0)
	return res
}
func dfsSubsets(res *[][]int, cur, nums []int, index int) {
	*res = append(*res, append([]int{}, cur...))
	for i := index; i < len(nums); i++ {
		cur = append(cur, nums[i])
		dfsSubsets(res, cur, nums, i+1)
		cur = cur[:len(cur)-1]

	}
}

// 最小覆盖子串
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	l, r, match, start, length := 0, 0, 0, 0, math.MaxInt32
	dict := map[byte]int{}
	for i := 0; i < len(t); i++ {
		dict[t[i]]++
	}
	for r < len(s) {
		if _, exist := dict[s[r]]; exist {
			dict[s[r]]--
			if dict[s[r]] == 0 {
				match++
			}
		}
		r++
		for match == len(dict) {
			if length > r-l+1 {
				length = r - l + 1
				start = l
			}
			if count, exist := dict[s[l]]; exist {
				if count == 0 {
					match--
				}
				dict[s[l]]++
			}
			l++
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : length+start]
}
