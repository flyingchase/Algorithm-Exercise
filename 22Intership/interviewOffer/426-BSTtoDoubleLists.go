package interviewoffer

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil {
		return root
	}
	if root.Left == nil && root.Right == nil {
		root.Left = root
		root.Right = root
		return root
	}
	stack := make([]*Node, 0)
	cur := root
	var prev, head *Node
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		// 首次遍历,记录头结点
		if prev == nil {
			head = node
		} else {
			// 非首次则prev 与出栈 node 链接
			prev.Right = node
			node.Left = prev
		}
		// prev 前驱记录当前node
		prev = node
		cur = node.Right
	}
	// 最后一个指针和 head 指针组成循环双链表
	if prev != nil && head != nil {
		prev.Right = head
		head.Left = prev
	}
	return head
}
