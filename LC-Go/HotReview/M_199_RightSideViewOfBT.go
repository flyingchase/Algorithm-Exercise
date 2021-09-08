package HotReview

import "LC-Go/DataStructure"

type TreeNode DataStructure.TreeNode
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

			//if queue[i].Left != nil {
			//	queue = append(queue, queue[i].Left)
			//}
			if queue[i].Right != nil {
				//queue = append(queue, queue[i].Right)
			}

			tmp = append(tmp, queue[i].Val)
		}
		queue = queue[size:]
		res = append(res, tmp[len(tmp)-1:]...)
	}
	return res
}
