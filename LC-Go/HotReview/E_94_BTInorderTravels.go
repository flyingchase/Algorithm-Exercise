package HotReview

//type TreeNode = DataStructure.TreeNode
func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res:=make([]int,0)
	// stack 为 TreeNode 的切片 📢 结构体需要指针
	stack :=make([]*TreeNode,0)
	cur:=root
	for cur!=nil||len(stack)>0 {
		// 一直向左侧走
		for cur!=nil {
			stack=append(stack,cur)
			cur=cur.Left
		}
		// 利用切片找到栈顶元素  弹出后注意 stack 需要丢掉栈顶
		node:=stack[len(stack)-1]
		stack=stack[:len(stack) - 1]
		res=append(res,node.Val)
		// 往上一层的右侧走 看右侧是否有左结点
		cur=node.Right
	}
	return res
}