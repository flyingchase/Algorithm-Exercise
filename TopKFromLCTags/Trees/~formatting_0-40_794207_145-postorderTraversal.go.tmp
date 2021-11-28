package trees

import (
	"topkTags"
)

type TreeNode = topkTags.TreeNode

func postorderTraversal(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := make([]int, 0)
	res = append(res, postorderTraversal(root.Left)...)
	res = append(res, postorderTraversal(root.Right)...)
	res = append(res, root.Val)
	return res

}
func post(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	cur, stack := root, make([]*TreeNode, 0)
	res := make([]int, 0)
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Left
	}
	for i := 0; i < len(res)>>1; i++ {
		res[i], res[len(res)-i-1] = res[len(res)-i-1], res[i]
	}
	return res
}
