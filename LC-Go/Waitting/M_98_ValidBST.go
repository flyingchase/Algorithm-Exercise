package Waitting

func isValidBST(root *TreeNode) bool {

	if root == nil {
		return true
	}
	ArrayBST := inOrderTraversalBST(root)

	for i:=0;i<len(ArrayBST)-1;i++ {
		if ArrayBST[i]>=ArrayBST[i+1] {
			return false
		}
	}

	return true
}

func inOrderTraversalBST(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack := make([]*TreeNode, 0)
	res := make([]int, 0)
	cur := root

	for cur != nil || len(stack) > 0 {
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
