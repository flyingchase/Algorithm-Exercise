package intership

// 递归前序遍历二叉树
func PreOrdreTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	res = append(res, root.Val)
	left := PreOrdreTraversalRecursion(root.Left)
	right := PreOrdreTraversalRecursion(root.Right)
	res = append(res, left...)
	res = append(res, right...)
	return res
}

// 递归中序遍历二叉树
func InOrdreTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	left := InOrdreTraversalRecursion(root.Left)
	right := InOrdreTraversalRecursion(root.Right)
	res = append(res, left...)
	res = append(res, root.Val)
	res = append(res, right...)
	return res
}

// 递归后序遍历二叉树
func PostOrdreTraversalRecursion(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}
	left := PostOrdreTraversalRecursion(root.Left)
	right := PostOrdreTraversalRecursion(root.Right)
	res = append(res, left...)
	res = append(res, right...)
	res = append(res, root.Val)
	return res
}
