package HotReview

import (
	"math"
)

/*
找到二叉树中的最大路径和

*/

// recursive
//type TreeNode=DataStructure.TreeNode

func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	getPathSum(root, &max)
	return max
}

func getPathSum(root *TreeNode, maxSum *int) int {
	if root == nil {
		return math.MaxInt32
	}
	left := getPathSum(root.Left, maxSum)
	right := getPathSum(root.Right, maxSum)
	curMax:= max(max(left+root.Val,right+root.Val),root.Val)
	*maxSum= max(*maxSum, max(curMax,left+root.Val+right))
	return curMax
}

func max(i int, j int) int {
	if i > j {
		return i

	}
	return j
}


func maxPathSum1(root *TreeNode) int {

	if root == nil {

		return 0

	}

	max := math.MinInt32

	getPathSum(root, &max)

	return max

}

func getPathSum1(root *TreeNode, maxSum *int) int {
	if root == nil { return math.MinInt32 }
	left := getPathSum(root.Left, maxSum)
	right := getPathSum(root.Right, maxSum)

	currMax := max(max(left+root.Val, right+root.Val), root.Val)
	*maxSum = max(*maxSum, max(currMax, left+right+root.Val))
	return currMax

}
