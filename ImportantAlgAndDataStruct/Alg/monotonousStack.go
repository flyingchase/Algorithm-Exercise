package alg

// 单调栈反向模板
// 仅保存数字（iterator point 为当前位置）
func nextGreaterElements1(nums []int) []int {
	n, res, stack := len(nums), make([]int, 0, len(nums)), make([]int, 0)
	for i := n - 1; i >= 0; i-- {
		for len(stack) != 0 && nums[i] >= stack[len(stack)-1] {
			stack = stack[:len(stack)-1]
		}
		if len(stack) == 0 {
			res[i] = -1
		} else {
			res[i] = stack[len(stack)-1]
		}
	}
	return res
}
