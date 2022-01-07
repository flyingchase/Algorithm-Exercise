package reviewlc60

func slidingWindow(nums []int, k int) []int {
	if len(nums) == 0 || len(nums) < k {
		return []int{}
	}
	windows := make([]int, 0, k)
	res := make([]int, 0, len(nums)-k)
	for i := 0; i < len(nums); i++ {
		// left index out bound of k
		for i >= k && windows[0] <= i-k {
			windows = windows[1:]
		}
		for len(windows) > 0 && nums[len(windows)-1] < nums[i] {
			windows = windows[:len(windows)-1]
		}
		windows = append(windows, i)
		if i > k {
			res = append(res, nums[windows[0]])
		}
	}
	return res
}
