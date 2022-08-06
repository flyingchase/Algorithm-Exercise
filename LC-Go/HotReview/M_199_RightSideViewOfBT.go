package HotReview

//type TreeNode=DataStructure.TreeNode

func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}

	cur := root

	queue := []*TreeNode{cur}
	res := make([]int, 0)

	for len(queue) > 0 {
		size := len(queue)
		tmp := make([]int, size)
		for i := 0; i < size; i++ {

			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}

			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[size:]
		res = append(res, tmp[len(tmp)-1:]...)
	}
	return res
}
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res := []int{}

	queue := make([]*TreeNode, 0)
	queue = append(queue, root)
	for len(queue) != 0 {
		size := len(queue)
		tmp := []int{}
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
			tmp = append(tmp, node.Val)
		}
		if len(tmp) != 0 {
			res = append(res, tmp[len(tmp)-1])
		}
	}

	return res
}
