package HotReview

import "math"

//import "LC-Go/DataStructure"

func isBalanced(root *TreeNode) bool {
	if root == nil {
		return true
	}
	leftHeigh := depth(root.Left)
	rigtHeight := depth(root.Right)
	return math.Abs(float64(leftHeigh-rigtHeight)) <= 1 && isBalanced(root.Left) && isBalanced(root.Right)

}

func depth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return maxisBalenced(depth(root.Right), depth(root.Left)) + 1
}

func maxisBalenced(i int, j int) int {
	if i > j {
		return i
	}
	return j

}
