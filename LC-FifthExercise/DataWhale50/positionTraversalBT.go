package datawhale50

import "sort"

func positionTraversalBT(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	type postitionTree struct {
		node     *TreeNode
		position int
	}
	// means 横坐标纵坐标当前节点值
	var levels map[int]map[int][]int
	queue := make([]postitionTree, 0)
	queue = append(queue, postitionTree{
		node:     root,
		position: 0,
	})

	deep := 0
	for len(queue) != 0 {
		size := len(queue)
		for i := 0; i < size; i++ {
			node := queue[0]

			queue = queue[1:]
			if node.node.Left != nil {
				queue = append(queue, postitionTree{
					node:     node.node.Left,
					position: node.position - 1,
				})
			}
			if node.node.Right != nil {
				queue = append(queue, postitionTree{
					node:     node.node.Right,
					position: node.position + 1,
				})
			}
			levelNode, ok := levels[node.position]
			if !ok {
				levelNode = make(map[int][]int)
			}
			levelNode[deep] = append(levelNode[deep], node.node.Val)
			levels[node.position] = levelNode
		}
		deep++
	}

	res := make([][]int, 0)
	for i := -1000; i < 1000; i++ {
		levelNode, ok := levels[i]
		if !ok {
			continue
		}
		levelVals := make([]int, 0)
		for j := 0; j <= deep; j++ {
			vals, ok := levelNode[j]
			if !ok {
				continue
			}
			sort.Ints(vals)
			levelVals = append(levelVals, vals...)
		}
		res = append(res, levelVals)
	}
	return res
}
