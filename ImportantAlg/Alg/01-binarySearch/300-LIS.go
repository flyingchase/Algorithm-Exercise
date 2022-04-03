package binarysearch

import "math"

// 最长递增子序列的长度

// dp[i]代表0~i 区间内最长递增子序列的长度
// 默认填充 1，本身一个单独的nums[i]也视为 lis
// O(N^2)
func lengthOfLIS_dp(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	dp := make([]int, len(nums))
	for index, _ := range dp {
		dp[index] = 1
	}
	res := 1

	for i := 1; i < len(nums); i++ {
		for j := 0; j < i; j++ {
			if nums[j] < nums[i] {
				dp[i] = int(math.Max(float64(dp[j]+1), float64(dp[i])))
			}
			res = int(math.Max(float64(res), float64(dp[i])))
		}
	}
	return res
}

// 构造 subSequence
// 但subSequence 不一定是结果，length 是相等的
// O(N^2)
// 查找第一个不比 num 大的数用二分则为 O(NlogN)
func lengthOfLIS_constSubSeqence(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sub := make([]int, 0)
	sub = append(sub, nums[0])
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > sub[len(sub)-1] {
			sub = append(sub, num)
		} else {
			// 找到第一个比 num 大或等的下标j，并更新
			j := 0
			for ; j < len(sub) && num > sub[j]; j++ {
			}
			sub[j] = num
		}
	}
	return len(sub)

}
func lengthOfLIS_SubSeqenceWithBinarySearch(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	sub := make([]int, 0)
	sub = append(sub, nums[0])
	for i := 1; i < len(nums); i++ {
		num := nums[i]
		if num > sub[len(sub)-1] {
			sub = append(sub, num)
		} else {
			// 找到第一个比 num 大或等的下标j，并更新
			j := binarysearchLIS(sub, num)
			sub[j] = num
		}
	}
	return len(sub)
}
func binarysearchLIS(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)>>1
		if nums[mid] == target {
			return mid
		} else if nums[mid] > target {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
