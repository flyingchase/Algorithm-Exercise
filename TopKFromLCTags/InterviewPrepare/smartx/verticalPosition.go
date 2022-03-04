package smartx

type TreeNode struct {
	Val         int
	Left, Right *TreeNode
}
type nodePosition struct {
	treeNode *TreeNode
	position int
}

func verticalTraversial(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	matrix, deep, queue := make(map[int]map[int][]int, 0), 0, make([]*nodePosition, 0)
	cur := root
	queue = append(queue, &nodePosition{
		treeNode: cur,
		position: 0,
	})
	for len(queue) > 0 {
		size := len(queue)
		for size > 0 {
			node := queue[0]
			queue = queue[1:]
			column, ok := matrix[node.position]
			if !ok {
				column = make(map[int][]int)
			}
			column[deep] = append(column[deep], node.treeNode.Val)
			if node.treeNode.Left != nil {
				queue = append(queue, &nodePosition{
					treeNode: node.treeNode.Left,
					position: node.position - 1,
				})
			}
			if node.treeNode.Right != nil {
				queue = append(queue, &nodePosition{
					treeNode: node.treeNode.Right,
					position: node.position + 1,
				})
			}
		}
		deep++
	}
	res := make([][]int, 0)
	for i := -1000; i < 1000; i++ {

	}
	return res

}
