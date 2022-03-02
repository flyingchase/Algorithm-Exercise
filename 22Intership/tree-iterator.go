package intership

// ========================================================================
func PreOrderTrvaersalBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, cur := make([]*TreeNode, 0), root
	res := []int{}
	for cur != nil || len(stack) > 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Right
	}
	return res
}

func InorderTraversalBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, cur, res := make([]*TreeNode, 0), root, []int{}
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			stack = append(stack, cur)
			cur = cur.Left
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		res = append(res, node.Val)
		cur = node.Right
	}
	return res
}

func PostorderTraversalBT(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res, cur, stack := []int{}, root, make([]*TreeNode, 0)
	for cur != nil || len(stack) != 0 {
		for cur != nil {
			res = append(res, cur.Val)
			stack = append(stack, cur)
			cur = cur.Right
		}
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		cur = node.Left
	}
	for i, j := 0, len(res)-1; i < j; i, j = i+1, j-1 {
		res[i], res[j] = res[j], res[i]
	}
	return res
}

// ========================================================================
func ZigzalTraversalaBT(root *TreeNode) [][]int {
	if root == nil {
		return nil
	}
	queue, res, cur := make([]*TreeNode, 0), [][]int{}, root
	queue = append(queue, cur)

	flag := false
	for len(queue) != 0 {
		size := len(queue)
		temp := []int{}
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			temp = append(temp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if flag {
			for i, j := 0, len(temp)-1; i < j; i, j = i+1, j-1 {
				temp[i], temp[j] = temp[j], temp[i]
			}
		}
		flag = !flag
		res = append(res, temp)
	}
	return res
}

// ========================================================================
// 前中后遍历构建二叉树
func ConstructBTFromPreInorderTraversal(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 {
		return nil
	}
	num := preorder[0]
	index := findIndexInTraversalRes(inorder, num)
	return &TreeNode{
		Val:   num,
		Left:  ConstructBTFromPreInorderTraversal(preorder[1:index+1], inorder[:index]),
		Right: ConstructBTFromPreInorderTraversal(preorder[index+1:], inorder[index+1:]),
	}
}
func findIndexInTraversalRes(nums []int, num int) int {
	var res int
	for i := 0; i < len(nums); i++ {
		if nums[i] == num {
			res = num
			break
		}
	}
	return res
}

func ConstructBTFromPostInorderTraversal(postorder []int, inorder []int) *TreeNode {
	if len(postorder) == 0 {
		return nil
	}
	num := postorder[len(postorder)-1]
	index := findIndexInTraversalRes(inorder, num)
	return &TreeNode{
		Val:   num,
		Left:  ConstructBTFromPostInorderTraversal(postorder[:index], inorder[:index]),
		Right: ConstructBTFromPostInorderTraversal(postorder[index:len(postorder)-1], inorder[index+1:]),
	}
}
