package datawhale50

import (
	"math"
	"sort"
)

// 最接近的三数之和
func threeSumClosest(nums []int, target int) int {
	// 先排序，保证后续双指针移动
	sort.Ints(nums)
	size := len(nums)
	// res 初始化为前三数之和,避免 size==3 时特殊情况
	res := nums[0] + nums[1] + nums[2]
	for i := 0; i < size-2; i++ {
		// l,r 双指针从 i 后到 size-1往中间移动
		l, r := i+1, size-1
		for l < r {
			sum := nums[i] + nums[l] + nums[r]
			// 注意 abs 判断, 更新 res 与 sum
			if math.Abs(float64(target-sum)) < math.Abs(float64(target-res)) {
				res = sum
			}
			// 双指针移动
			if sum <= target {
				l++
			} else {
				r--
			}
		}
	}
	return res
}
