package interviewprepare

import "math"

func maxPathSum(root *TreeNode) int {
	maxSum := math.MinInt32
	maxTree(root, &maxSum)
	return maxSum
}
func maxTree(node *TreeNode, maxSum *int) int {
	if node == nil {
		return 0
	}
	left := maxMaxPathSum(maxTree(node.Left, maxSum), 0)
	right := maxMaxPathSum(maxTree(node.Right, maxSum), 0)
	*maxSum = maxMaxPathSum(*maxSum, node.Val+left+right)
	return node.Val + maxMaxPathSum(left, right)
}
func maxMaxPathSum(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
