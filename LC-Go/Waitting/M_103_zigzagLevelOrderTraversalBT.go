package Waitting

func zigzagLevelOrder(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}

	res := make([][]int, 0)
	cur := root

	queue := make([]*TreeNode, 1)
	queue = append(queue, cur)
	flag := false

	for len(queue) > 0 {
		size := len(queue)
		lists := make([]int, 0)
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			if node == nil {
				continue
			}
			lists = append(lists, node.Val)
			queue = append(queue, node.Left)
			queue = append(queue, node.Right)
		}
		if flag {
			for i := 0; i < len(lists)>>1; i++ {
				lists[i], lists[len(lists)-1-i] = lists[len(lists)-i-1], lists[i]
			}
		}
		flag = !flag
		if len(lists) > 0 {
			res = append(res, lists)
		}
	}
	return res
}
