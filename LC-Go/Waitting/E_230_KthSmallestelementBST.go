package Waitting

import "LC-Go/DataStructure"

type TreeNode = DataStructure.TreeNode
func kthSmallest(root *TreeNode, k int) int {
	if root == nil {
		return -1
	}


	stack:=make([]*TreeNode,0)
	res:=make([]int,0)
	cur:=root
	for cur!=nil||len(stack) >0{
		for cur!=nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node:=stack[len(stack) - 1]
		stack=stack[:len(stack)-1]
		res = append(res, node.Val)
		cur=node.Right
	}
	return res[k-1]
}


