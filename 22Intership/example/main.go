package main

import (
	intership "22intership"
	"fmt"
)

type TreeNode = intership.TreeNode

func main() {
	// 1,2,3,4,5,6
	root := &TreeNode{
		Val: 1,
		Left: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val:   4,
				Left:  nil,
				Right: nil,
			},
			Right: &TreeNode{
				Val:   5,
				Left:  nil,
				Right: nil,
			},
		},
		Right: &TreeNode{
			Val: 3,
			Left: &TreeNode{
				Val:   6,
				Left:  nil,
				Right: nil,
			},
			Right: nil,
		},
	}

	fmt.Println(intership.PreOrderTrvaersalBT(root))

	matrix := [][]int{{0, 0, 0, 0}, {1, 0, 1, 0}, {0, 0, 0, 0}, {1, 2, 3, 4}, {1, 2, 3, 4}}
	fmt.Println(len(matrix[0]))
}
