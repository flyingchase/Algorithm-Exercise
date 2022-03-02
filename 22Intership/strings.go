package intership

import (
	"math"
	"reflect"
)

// ========================================================================
// 最小子数组之和>target的长度
func minSubArrayLen(nums []int, target int) int {
	res := math.MaxInt32
	l, r, sum := 0, 0, 0
	// 右指针前进，不断累积右指针
	for r < len(nums) {
		sum += nums[r]
		// 符合条件则收缩
		for sum >= target {
			// 更新 res
			if r-l+1 < res {
				res = r - l + 1
			}
			// shrink
			sum -= nums[l]
			l++
		}
		r++
	}
	// 特例判断
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

// ========================================================================
// 至多有 k 个不同字符的最长子串
func LengthOfLongestSubstringKDistinct(s string, k int) int {
	dict := make(map[byte]int, 0)
	l, r, res := 0, 0, 0
	for r < len(s) {
		// 右指针前进，不断计数出现的字符的次数
		dict[byte(s[r])]++
		// 出现字符个数符合条件，收缩
		for len(dict) > k {
			// shrink
			dict[byte(s[l])]--
			if dict[byte(s[l])] == 0 {
				delete(dict, byte(s[l]))
			}
			l++
		}
		if r-l+1 > res {
			res = r - l + 1
		}
		r++
	}
	return res
}

// ========================================================================
// 判断给定字符串 s 是否含有 p 的排列
// 定长滑动窗口，两个 map 分别记录 p 中字符次数和窗口中字符频次
func checkInclusion(s1, s2 string) bool {
	len1, len2 := len(s1), len(s2)
	if len1 > len2 {
		return false
	}
	map1, map2 := make([]int, 26), make([]int, 26)
	// 首个窗口入 map2，同时记录 s1中的字符频次
	for i := 0; i < len1; i++ {
		map1[s1[i]-'a']++
		map2[s2[i]-'a']++
	}
	// 在字符串s2 上从 len1到 len2 滑动窗口
	// 窗口大小固定为 len1
	for i := len1; i < len2; i++ {
		if reflect.DeepEqual(map2, map1) {
			return true
		}
		// 加入新的字符
		map2[s2[i]-'a']++
		// 移出 i-len1 窗口最左侧的字符频次
		map2[s2[i-len1]-'a']--
	}
	return reflect.DeepEqual(map1, map2)
}

// ========================================================================
// 最小子串覆盖
// s 中包含 t 中每个字符的子串的最小子串
func MinWindow_76(s, t string) string {
	dict, start, l, r := make(map[byte]int, 0), 0, 0, 0
	length := math.MaxInt32
	match := 0
	// dict 存储 target 中字符频次
	for i := 0; i < len(t); i++ {
		dict[t[i]]++
	}
	// 在 s 上滑动窗口
	for r < len(s) {
		// 遇见s[r]符合条件
		// 在 dict 内，则将字符频次--
		if _, boolV := dict[s[r]]; boolV {
			dict[s[r]]--
			// 改字符在目前 l-r 内已经被匹配
			if dict[s[r]] == 0 {
				match++
			}
		}
		// 所有字符均在 s[l:r+1]内被匹配
		for match == len(dict) {
			// 更新有效子串的长度，同时记录起始位置
			if length > r-l+1 {
				length = r - l + 1
				start = l
			}
			// shrink left window boarder
			if count, boolValue := dict[s[l]]; boolValue {
				// 更新已经匹配的字符
				if count == 0 {
					match--
				}
				dict[s[l]]++
			}
			l++
		}
		r++
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}
