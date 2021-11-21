package exercise

type TreeNode struct {
	Left  *TreeNode
	Right *TreeNode
	Val   int
}

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	queue := make([]*TreeNode, 0)
	cur := root
	queue = append(queue, cur)

	res := make([]int, 0)
	if len(queue) != 0 {
		lists := make([]int, 0)
		size := len(queue)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			lists = append(lists, node.Val)
			if node.Left != nil {
				node = node.Left
				queue = append(queue, node)
			}
			if node.Right != nil {
				node = node.Right
				queue = append(queue, node)
			}
		}
		if len(lists) != 0 {
			res = append(res, lists[len(lists)-1])
		}
	}
	return res
}
func cloestAncessorInBT(root *TreeNode, p, q *TreeNode) *TreeNode {
	if root == nil || root == q || root == p {
		return root
	}
	left := cloestAncessorInBT(root.Left, p, q)
	right := cloestAncessorInBT(root.Right, p, q)
	if left != nil {
		if right != nil {
			return root
		}
		return left
	}
	return right
}
func reverseBetween(head *ListNode, m, n int) *ListNode {
	if head == nil || m >= n {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	cur := dummyHead
	for i := 0; i < m-1; i++ {
		cur = cur.Next
	}
	wall := cur
	leftNode := wall.Next
	rightNode := leftNode.Next
	for i := m + 1; i <= n; i++ {
		leftNode.Next = rightNode.Next
		rightNode.Next = wall.Next
		wall.Next = rightNode
		rightNode = leftNode.Next
	}
	return dummyHead.Next
}
