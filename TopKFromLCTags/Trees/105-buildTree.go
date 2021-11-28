package trees

// recursive 建造二叉树
func buildTree1(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	num := preorder[0]
	index := findIndex(inorder, num)
	return &TreeNode{
		Val:   num,
		Left:  buildTree1(preorder[1:index+1], inorder[:index]),
		Right: buildTree1(preorder[index+1:], inorder[index+1:]),
	}
}
func findIndex(inorder []int, target int) int {
	var index int
	for i, num := range inorder {
		if target == num {
			index = i
			break
		}
	}
	return index
}
