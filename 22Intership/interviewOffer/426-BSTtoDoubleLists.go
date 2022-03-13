package interviewoffer

type Node struct {
	Val   int
	Left  *Node
	Right *Node
}

func treeToDoublyList(root *Node) *Node {
	if root == nil || (root.Left == nil && root.Right == nil) {
		return root
	}
	stack := make([]*Node, 0)
	cur := root
	dummtRoot := &Node{Val: -1}
	var min *Node
	first := true
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		if first {
			min = node
			first = false
			dummtRoot.Left = min
		}
		stack = stack[:len(stack)-1]
		cur = node.Right
		node.Left = cur
	}
	return min
}
