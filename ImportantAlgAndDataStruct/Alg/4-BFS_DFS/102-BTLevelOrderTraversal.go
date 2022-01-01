package bfsdfs

import (
	"alg"
)

type TreeNode = alg.TreeNode

// M-102 二叉树层序遍历
func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, cur := make([]*TreeNode, 0), root
	res := [][]int{}
	queue = append(queue, cur)
	for len(queue) != 0 {
		temp := make([]int, 0)
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			temp = append(temp, node.Val)
		}
		res = append(res, temp)
	}
	return res
}
