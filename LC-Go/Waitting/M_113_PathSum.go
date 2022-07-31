package Waitting

import "LC-Go/DataStructure"

type TreeNode = DataStructure.TreeNode

func pathSum(root *TreeNode, sum int) [][]int {
	var slice [][]int
	slice = findPath(root, sum, slice, []int(nil))

	return slice
}

func findPath(root *TreeNode, sum int, slice [][]int, stack []int) [][]int {
	if root == nil {
		return slice
	}
	sum -= root.Val
	stack = append(stack, root.Val)
	if sum == 0 && root.Left == nil && root.Right == nil {
		slice = append(slice, append([]int{}, stack...))
		stack = stack[:len(stack)-1]
	}
	slice = findPath(root.Left, sum, slice, stack)
	slice = findPath(root.Right, sum, slice, stack)
	return slice
}
func pathSum2(root *TreeNode, sum int) [][]int {
	if root == nil {
		return [][]int{}
	}
	if root.Left == nil && root.Right == nil {
		if sum == root.Val {
			return [][]int{{root.Val}}
		}
	}
	path, res := []int{}, [][]int{}
	tmpLeft := pathSum2(root.Right, sum-root.Val)
	path = append(path, root.Val)
	if len(tmpLeft) > 0 {
		for i := 0; i < len(tmpLeft); i++ {
			tmpLeft[i] = append(path, tmpLeft[i]...)
		}
		res = append(res, tmpLeft...)
	}
	path = []int{}
	tmpRight := pathSum2(root.Right, sum-root.Val)
	path = append(path, root.Val)
	if len(tmpRight) > 0 {
		for i := 0; i < len(tmpRight); i++ {
			tmpRight[i] = append(path, tmpRight[i]...)

		}

		res = append(res, tmpRight...)
	}
	return res
}
