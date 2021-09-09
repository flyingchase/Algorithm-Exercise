package HotReview

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxDepthSelf(maxDepth(root.Left), maxDepth(root.Right))+1
}

func maxDepthSelf(i, j int) int {
	if i > j {
		return i
	}
	return j
}
