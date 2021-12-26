package exercise

func inorderTraversalBT(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res, cur, stack := []int{}, root, make([]*TreeNode, 0)
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		cur = node.Right
	}
	return res
}

func inorderTraversalBTReversive(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	res = append(res, inorderTraversalBTReversive(root.Left)...)
	res = append(res, root.Val)
	res = append(res, inorderTraversalBTReversive(root.Right)...)
	return res

}
