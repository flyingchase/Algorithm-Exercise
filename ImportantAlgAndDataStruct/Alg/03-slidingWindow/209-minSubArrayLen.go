package slidingwindow

import "math"

// 209-长度最小的子数组
// 找到连续子数组的和>=target且长度最小
// 返回符合的最小连续子数组的长度
func minSubArrayLen(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, 0
	// 初始 res 很大，使得res>right-left+1首次时显然成立
	// 否则初始为 0 则不会更新 res
	res := math.MaxInt32
	sum := 0
	for right < len(nums) {
		sum += nums[right]
		// slidingwindow 一般可以用 for 来 shrink 缩小窗口
		for sum >= target {
			if res > right-left+1 {
				res = right - left + 1
			}
			sum -= nums[left]
			left++
		}
		right++
	}
	// 判断 res 是否被更新过，未更新则找不到>=target的子数组，返回 0
	if res == math.MaxInt32 {
		return 0
	}
	return res
}

// 二分猜答案
// 上下界为 0 len-1
// 符合子数组之和>=target即为 valid size
func minSubArrayLen_binarySearch(target int, nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	left, right := 0, len(nums)-1
	res := 0
	for left <= right {
		mid := left + (right-left)>>1
		if valid(nums, mid, target) {
			right = mid - 1
			res = mid
		} else {
			left = mid + 1
		}
	}
	return res
}

func valid(nums []int, size, target int) bool {
	sum := 0
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		// 保证sum 累加的个数不超过 size
		if i-size >= 0 {
			sum -= nums[i-size]
		}
		if sum >= target {
			return true
		}
	}
	return false
}
