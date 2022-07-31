package Waitting

func rebuildBT(preorder, inorder []int) *TreeNode {
	if len(preorder)*len(inorder) == 0 {
		return nil
	}
	index := findIndex(preorder[0], inorder)
	return &TreeNode{
		Val:   preorder[0],
		Left:  rebuildBT(preorder[1:index+1], inorder[:index]),
		Right: rebuildBT(preorder[index+1:], inorder[index+1:]),
	}
}
func findIndex(num int, nums []int) int {
	for index, value := range nums {
		if value == num {
			return index
		}
	}
	return 0
}
