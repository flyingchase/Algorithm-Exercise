package reviewlc60

import (
	"alg"
)

type TreeNode = alg.TreeNode

func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == p || root == q {
		return root
	}
	left := lca(root.Left, p, q)
	right := lca(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
