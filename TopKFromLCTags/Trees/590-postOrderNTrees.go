package trees

type Node struct {
	Val      int
	Children []*Node
}

func postorder(root *Node) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	for _, item := range root.Children {
		res = append(res, postorder(item)...)
	}
	res = append(res, root.Val)
	return res
}

func postorderIterator(root *Node) []int {
	if root == nil {
		return []int{}
	}

	cur, stack := root, make([]*Node, 0)
	res := make([]int, 0)
	stack = append(stack, cur)
	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		stack = append(stack, node.Children...)
	}
	for i := 0; i < len(res)>>1; i++ {
		res[i], res[len(res)-1-i] = res[len(res)-i-1], res[i]
	}
	return res

}
