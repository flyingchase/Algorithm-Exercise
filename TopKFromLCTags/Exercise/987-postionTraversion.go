package exercise

import "sort"

type nodePositon struct {
	node     *TreeNode
	position int
}

// 二叉树的垂直遍历
func verticalTraversal(root *TreeNode) [][]int {

	if root == nil {
		return [][]int{}
	}
	matrix := make(map[int]map[int][]int)
	queue := make([]*nodePositon, 0)
	queue = append(queue, &nodePositon{node: root, position: 0})
	deep := 0
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			curNode := queue[0]
			queue = queue[1:]
			column, ok := matrix[curNode.position]
			if !ok {
				column = make(map[int][]int)
			}
			column[deep] = append(column[deep], curNode.node.Val)
			matrix[curNode.position] = column
			if curNode.node.Left != nil {
				queue = append(queue, &nodePositon{node: curNode.node.Left, position: curNode.position - 1})
			}
			if curNode.node.Right != nil {
				queue = append(queue, &nodePositon{node: curNode.node.Right, position: curNode.position + 1})
			}
		}
		deep++
	}

	res := [][]int{}

	for position := -1000; position <= 1000; position++ {
		column, ok := matrix[position]
		if !ok {
			continue
		}
		vals := make([]int, 0)
		for curDeep := 0; curDeep < deep; curDeep++ {
			nodes, ok := column[curDeep]
			if !ok {
				continue
			}
			sort.Ints(nodes)
			vals = append(vals, nodes...)
		}
		res = append(res, vals)
	}
	return res
}
