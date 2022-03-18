package backtrack

import (
	"suanfa8"
)

type TreeNode = suanfa8.TreeNode

func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	backtrackPathSum(root, &res, []int{}, targetSum)
	return res
}
func backtrackPathSum(node *TreeNode, res *[][]int, cur []int, carry int) {
	cur = append(cur, node.Val)
	if node.Left == nil && node.Right == nil && carry == 0 {
		*res = append(*res, append([]int{}, cur...))
		return
	}
	if node.Left != nil {
		backtrackPathSum(node.Left, res, cur, carry-node.Val)
	}
	if node.Right != nil {
		backtrackPathSum(node.Right, res, cur, carry-node.Val)
	}
	cur = cur[:len(cur)-1]
}
