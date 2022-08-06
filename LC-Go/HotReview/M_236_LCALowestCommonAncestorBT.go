package HotReview

/*
二叉树中指定两节点的最近公共祖先结点
*/
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}

	left := lowestCommonAncestor(root.Left, p, q)
	right := lowestCommonAncestor(root.Right, p, q)

	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right

}

func lowestCommonAncestor2(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	l, r := lowestCommonAncestor2(root.Left, p, q), lowestCommonAncestor2(root.Right, p, q)
	if l != nil {
		if r != nil {
			return root
		}
		return l
	}
	return r
}
