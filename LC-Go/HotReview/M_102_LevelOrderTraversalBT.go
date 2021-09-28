package HotReview



func levelOrder(root *TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}
	
	queue := []*TreeNode{root}
	
	res := make([][]int, 0)
	for len(queue) > 0 {
		l := len(queue)
		temp := make([]int, 0, l)
		for i := 0; i < l; i++ {
			if queue[i].Left != nil {
				queue = append(queue, queue[i].Left)
			}
			if queue[i].Right != nil {
				queue = append(queue, queue[i].Right)
			}
			temp = append(temp, queue[i].Val)
	
		}
		queue = queue[l:]
		res = append(res, temp)
	}
	return res
}
