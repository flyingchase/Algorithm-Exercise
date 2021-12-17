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
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
	}
}
func findIndexInPostOrder(nums []int, target int) int {
	var index int
	for i, num := range nums {
		if target == num {
			index = i
			break
		}
	}
	return index
}
