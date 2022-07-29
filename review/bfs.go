package review

func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, res, cur := []*TreeNode{}, [][]int{}, root
	queue = append(queue, cur)
	for len(queue) != 0 {
		size := len(queue)
		tmp := []int{}
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
		if len(tmp) != 0 {
			res = append(res, tmp)
		}
	}
	return res
}
