package bytedance

/*
*给定二叉树中的两节点，求两节点之间的最短路径
 1. 求两节点之间的LCA，最近公共祖先
 2. 求公共祖先到两节点之间的距离
 3. 累加两距离
* */

func findShortestPathInTwoNodesInBT(root, p, q *TreeNode) int {
	if p == nil || q == nil {
		return -1
	}
	lcaNode := lca(root, p, q)
	len1, len2 := getPath(lcaNode, p), getPath(lcaNode, q)
	return len1 + len2
}
func lca(root, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left, right := lca(root.Left, p, q), lca(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
func getPath(lcaNode, targetNode *TreeNode) int {
	if lcaNode == nil {
		return 0
	}
	left, right := getPath(lcaNode.Left, targetNode), getPath(lcaNode.Right, targetNode)
	if lcaNode == targetNode {
		return 1
	}
	if left == 0 && right == 0 {
		return 0
	}
	if left == 0 {
		return right + 1
	}
	return left + 1
}
