package HotReview

/*
对称二叉树
*/
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	return judge(left, right)
}

func judge(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}
	return judge(left.Right, right.Left) && judge(left.Left, right.Right)
}
