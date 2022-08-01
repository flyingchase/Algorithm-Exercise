package SetSailAgain

// 验证后序遍历序列是否为 BST 树
// 左<根<右
func verifyPostorder(postorder []int) bool {
	if len(postorder) <= 1 {
		return true
	}
	rootVal := postorder[len(postorder)-1]
	leftBound := 0

	for k, v := range postorder {
		if v >= rootVal {
			leftBound = k
			break
		}
	}
	left, right := postorder[:leftBound], postorder[leftBound:len(postorder)-1]
	// 依据bst性质验证
	// leftpart 已经在查找 index 时候验证过了
	// for _, v := range left {
	// 	if v > rootVal {
	// 		return false
	// 	}
	// }
	for _, v := range right {
		if v < rootVal {
			return false
		}
	}
	return verifyPostorder(left) && verifyPostorder(right)
}
