package interviewprepare

func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	num := preorder[0]
	index := findIndex(inorder, num)
	leftNode := buildTree(preorder[1:index+1], inorder[:index])
	rightNode := buildTree(preorder[index+1:], inorder[index+1:])
	return &TreeNode{
		Val:  num,
		Left: leftNode, Right: rightNode,
	}
}
func buildTreeFromInAndPostTraversalRes(inorder []int, postorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	num := postorder[len(postorder)-1]
	index := findIndex(inorder, num)
	return &TreeNode{
		Val:   num,
		Left:  buildTree(inorder[:index], postorder[:index]),
		Right: buildTree(inorder[index+1:], postorder[index:len(postorder)-1]),
	}

}
func findIndex(nums []int, target int) int {
	var index int
	for i, num := range nums {
		if num == target {
			index = i
			break
		}
	}
	return index
}
