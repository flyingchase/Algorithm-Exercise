package trees

func zigzagLevelOrder(root *TreeNode) [][]int {

	if root == nil {
		return nil
	}
	res, queue, flag, cur := [][]int{}, make([]*TreeNode, 0), false, root
	queue = append(queue, cur)
	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, 0)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if flag {
			reverse(tmp)
		}
		flag = !flag
		res = append(res, tmp)
	}
	return res
}
func reverse(nums []int) {
	for i := 0; i < len(nums)>>1; i++ {
		nums[i], nums[len(nums)-i-1] = nums[len(nums)-1-i], nums[i]
	}
}
