package HotReview

import (
	"fmt"
)

/*
对称二叉树
*/

//type TreeNode = DataStructure.TreeNode
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return true
	}
	left := root.Left
	right := root.Right
	fmt.Println("this is writed by vim\n")
	nums := make([]int, 1)
	nums = append(nums, 1)
	for k, v := range nums {
		fmt.Println(k, v)
	}
	if len(nums) > 0 {
		fmt.Println(len(nums))

	}
	fmt.Println(nums)
	return judge(left, right)
}

func judge(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil {
		return false
	}
	if left.Val != right.Val {
		return false
	}

	return judge(left.Right, right.Left) && judge(left.Left, right.Right)
}
