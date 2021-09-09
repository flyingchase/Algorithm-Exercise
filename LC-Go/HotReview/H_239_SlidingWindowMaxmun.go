package HotReview

/*
滑动窗口内的最大值
*/

// sloution1 暴力 NK
func maxSlidngWindow1(nums []int, k int) []int {
	res := make([]int, 0, k)
	n := len(nums)
	if n == 0 {
		return []int{}
	}

	for i := 0; i <= n-k; i++ {
		max := nums[i]
		// form i kStep
		for j := 1; j < k; j++ {
			if max < nums[i+j] {
				max = nums[i+j]
			}
		}
		res = append(res, max)
	}
	return res

}

// Solution  优先队列

// Solution3 双端队列 O(N)   O(K)

func maxSlidingWindows(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return make([]int, 0)
	}
	// window store the index of nums
	window := make([]int, 0, k)
	res := make([]int, 0, len(nums)-k+1)
	for i, v := range nums {
	// left index  out of the window remove
	//	i>=k means i not the scale of first window
		if i >= k && window[0] <= i-k {
			window = window[1:len(window)]
		}
		// maintain window
		// the right of window in nums is bigger the curCycle remove the right
		for len(window) > 0 && nums[window[len(window)-1]] < v {
			window = window[0 : len(window)-1]
		}
		window = append(window, i)
		// 保证窗口大小在 k 以上的时候再存储 windown内左侧的 index（对应 nums 内该区间的最值）
		if i >= k-1 {
			res = append(res, nums[window[0]])

		}
	}
	return res
}
