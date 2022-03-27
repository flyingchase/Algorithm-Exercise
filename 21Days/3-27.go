package days

// 二叉树的直径
func diameterOfBinaryTree(root *TreeNode) int {
	res := 0
	helperDia(root, &res)
	return res
}
func helperDia(node *TreeNode, res *int) int {
	if node == nil {
		return 0
	}
	left, right := helperDia(node.Left, res), helperDia(node.Right, res)
	if left+right > *res {
		*res = left + right
	}
	*res = max(left+right, *res)
	return max(left, right) + 1
}

// 路径总和
func hasPathSum(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	if root.Left == nil && root.Right == nil {
		return root.Val == targetSum
	}
	return hasPathSum(root.Left, targetSum-root.Val) || hasPathSum(root.Right, targetSum-root.Val)
}

// 路径总和 回溯非精简代码
func hasPathSum2(root *TreeNode, targetSum int) bool {
	if root == nil {
		return false
	}
	res := false
	dfsHasPathSum(root, targetSum, 0, &res)
	return res
}
func dfsHasPathSum(node *TreeNode, targetSum, cur int, res *bool) {
	if cur == targetSum && node.Left == nil && node.Right == nil {
		*res = true
		return
	}
	if node.Left != nil {
		cur += node.Left.Val
		dfsHasPathSum(node.Left, targetSum, cur, res)
		cur -= node.Left.Val
	}
	if node.Right != nil {
		cur += node.Right.Val
		dfsHasPathSum(node.Right, targetSum, cur, res)
		cur -= node.Right.Val
	}
}

// 平衡二叉树
// 自顶向下依次遍历
func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	if maxBTDepth(root.Left)-maxBTDepth(root.Right) > 1 || maxBTDepth(root.Left)-maxBTDepth(root.Right) < -1 {
		return false
	}
	return isBalanced(root.Left) && isBalanced(root.Right)
}
func maxBTDepth(node *TreeNode) int {
	if node != nil {
		l, r := maxBTDepth(node.Left), maxBTDepth(node.Right)
		if l > r {
			return l + 1
		} else {
			return r + 1
		}
	}
	return 0
}

// 根节点到叶子结点路径的总和
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	}
	res := 0
	dfsSumNumbers(&res, root, 0)
	return res
}
func dfsSumNumbers(res *int, node *TreeNode, cur int) {
	if node == nil {
		return
	}
	cur = cur*10 + node.Val
	if node.Left == nil && node.Right == nil {
		*res += cur
		return
	}
	if node.Left != nil {
		dfsSumNumbers(res, node.Left, cur)
	}
	if node.Right != nil {
		dfsSumNumbers(res, node.Right, cur)
	}
}
