package exercise

func recursiveTraversalBT_posrOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	res = append(res, recursiveTraversalBT_posrOrder(root.Left)...)
	res = append(res, recursiveTraversalBT_posrOrder(root.Right)...)
	res = append(res, root.Val)
	return res
}

func TraversalBT_posrOrder(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	stack, cur := make([]*TreeNode, 0), root
	res := []int{}
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Left
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

