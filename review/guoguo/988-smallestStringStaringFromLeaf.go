package guoguo

import (
	"review"
	"sort"
)

type TreeNode review.TreeNode

// var res []string

func smallestFromLeaf(root *TreeNode) string {
	var path string
	if root == nil {
		return path
	}
	preorderSmallest(root, &path)
	tmp := []byte(path)
	sort.Slice(tmp, func(i, j int) bool {
		if tmp[i] < tmp[j] {
			return true
		}
		return false
	})
	path = string(tmp)
	return path
}

func preorderSmallest(root *TreeNode, res *string) {
	if root == nil {
		return
	}
	ch := byte(root.Val) + byte('a')
	(*res) += string(ch)
	if root.Left != nil {
		preorderSmallest((*TreeNode)(root.Left), res)
	}
	if root.Right != nil {
		preorderSmallest((*TreeNode)(root.Right), res)
	}
	if root.Left == nil && root.Right == nil {
		return
	}
}
