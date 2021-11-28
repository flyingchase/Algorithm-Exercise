package trees

func buildTree(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	num := postorder[len(postorder)-1]
	index := findIndexInPostOrder(inorder, num)
	return &TreeNode{
		Val:   num,
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index+1:len(postorder)-1]),
	}
}
func findIndexInPostOrder(nums []int, target int) int {
	for i, num := range nums {
		if target == num {
			return i
		}
	}

	return -1
}
