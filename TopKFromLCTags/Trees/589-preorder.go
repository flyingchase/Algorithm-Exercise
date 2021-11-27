package trees

func preorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, root.Val)
	for _, item := range root.Children {
		res = append(res, preorder(item)...)
	}
	return res
}
func preorderIterator(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	cur := root
	stack := make([]*Node, 0)
	stack = append(stack, cur)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		for i := len(node.Children) - 1; i >= 0; i-- {
			stack = append(stack, node.Children[i])
		}
	}
	return res
}
