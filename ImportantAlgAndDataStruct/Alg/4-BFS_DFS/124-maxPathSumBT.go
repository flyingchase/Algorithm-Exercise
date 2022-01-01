package bfsdfs

import (
	"math"
)

/*
返回二叉树中的最大路径和
任意两个结点间的序列即为路径，路径中各节点值的和即为路径和
不一定经过 root
*/

// 最大路径和回溯 recursion
// O(N)
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	helper(root, &max)
	return max
}
func helper(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	left := int(math.Max(float64(0), float64(helper(node.Left, max))))
	right := int(math.Max(float64(0), float64(helper(node.Right, max))))
	// 横向比较
	*max = int(math.Max(float64(*max), float64(left+right+node.Val)))
	// 纵向比较
	return int(math.Max(float64(left), float64(right))) + node.Val
}
