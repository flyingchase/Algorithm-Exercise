package SetSailAgain

import (
	"LC-Go/DataStructure"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */

type TreeNode = DataStructure.TreeNode

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(inorder) == 0 || len(preorder) == 0 {
		return nil
	}
	index := findIndex(inorder, preorder[0])
	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
}

func findIndex(nums []int, target int) int {

	for i := 0; i < len(nums); i++ {
		if nums[i] == target {
			return i
		}
	}
	return 0
}
