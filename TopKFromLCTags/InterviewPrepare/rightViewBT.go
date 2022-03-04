package interviewprepare

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}

func rightView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	queue, cur := make([]*TreeNode, 0), root
	queue = append(queue, cur)
	res := []int{}
	for len(queue) > 0 {
		size := len(queue)
		level := make([]int, 0)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			level = append(level, node.Val)
		}
		res = append(res, level[len(level)-1])
	}
	return res
}
