package exercise

func bstToLists(root *TreeNode) *ListNode {
	if root == nil {
		return nil
	}
	cur, dummyHead := root, &ListNode{Val: -1}
	curNode := dummyHead
	stack := make([]*TreeNode, 0)
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		curNode.Next = &ListNode{Val: node.Val}
		cur = node.Right
	}
	return dummyHead.Next
}
