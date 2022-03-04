package interviewprepare

func pathSum(root *TreeNode, targetSum int) [][]int {
	var res [][]int
	if root == nil {
		return res
	}
	helper(root, []int{}, targetSum, &res)
	return res
}
func helper(root *TreeNode, path []int, targetSum int, res *[][]int) {
	if root.Left == nil && root.Right == nil {
		if root.Val == targetSum {
			*res = append(*res, append(path, root.Val))
		}
		return
	}
	if root.Left != nil {
		helper(root.Left, append(append([]int{}, path...), root.Val), targetSum-root.Val, res)
	}
	if root.Right != nil {
		helper(root.Right, append(append([]int{}, path...), root.Val), targetSum-root.Val, res)
	}
}
