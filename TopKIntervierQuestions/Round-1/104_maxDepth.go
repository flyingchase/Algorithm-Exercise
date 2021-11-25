package round1

import (
	"math"
	"topk"
)

type TreeNode = topk.TreeNode

// dfs
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	cur, queue := root, make([]*TreeNode, 0)
	queue = append(queue, cur)
	res := 0
	for len(queue) > 0 {
		size := len(queue)
		if size > 0 {
			res++
		}
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
		}
	}
	return res
}

// recursive
func maxDepth2(root *TreeNode) int {
	if root == nil {
		return 0
	}
	count:=1+math.Max(float64(maxDepth2(root.Left)),float64(maxDepth2(root.Right)))
	return int(count)
	// leftDepth := maxDepth2(root.Left)
	// rightDepth := maxDepth2(root.Right)
	// if leftDepth > rightDepth {
	// 	return leftDepth + 1
	// } else {
	// 	return rightDepth + 1
	// }
}
