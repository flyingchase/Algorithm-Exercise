package trees

import "sort"

type nodePositon struct {
	node     *TreeNode
	position int
}

func verticalTraversal(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	matrix, deep := make(map[int]map[int][]int), 0
	queue := make([]*nodePositon, 0)
	queue = append(queue, &nodePositon{node: root, position: 0})
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			size--
			curNode := queue[0]
			queue = queue[1:]
			coloum, ok := matrix[curNode.position]
			if !ok {
				coloum = make(map[int][]int)
			}
			coloum[deep] = append(coloum[deep], curNode.node.Val)

			matrix[curNode.position] = coloum
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
	for curPosition := -1000; curPosition <= 1000; curPosition++ {
		curColumn, ok := matrix[curPosition]
		if !ok {
			continue
		}
		vals := make([]int, 0)
		for curDeep := 0; curDeep < deep; curDeep++ {
			curLevels, ok := curColumn[curDeep]
			if !ok {
				continue
			}
			sort.Ints(curLevels)
			vals = append(vals, curLevels...)
		}
		res = append(res, vals)
	}
	return res
}
