package trees

func levelOrderBottom(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	res := levelOrder107(root)
	for i := 0; i < len(res)>>1; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}

	return res
}

func levelOrder107(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	queue, res, cur := make([]*TreeNode, 0), [][]int{}, root
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
		res = append(res, tmp)
	}
	return res
}
