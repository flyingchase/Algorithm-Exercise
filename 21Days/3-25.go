package days

import (
	"strconv"
	"strings"
)

func reverseWords(s string) string {
	if len(s) == 0 {
		return s
	}
	index := 0
	for ; index < len(s) && s[index] == ' '; index++ {
	}
	s = s[index:]
	tail := len(s) - 1
	for ; tail >= 0 && s[tail] == ' '; tail-- {
	}
	s = s[:tail+1]
	str := []byte(s)
	for i := 1; i < len(str); i++ {
		if str[i] == ' ' && str[i-1] == ' ' {
			str = append(str[:i], str[i+1:]...)
			i--
		}
	}
	ss := strings.Split(string(str), " ")
	for i, j := 0, len(ss)-1; i < j; i, j = i+1, j-1 {
		ss[i], ss[j] = ss[j], ss[i]
	}
	return strings.Join(ss, " ")
}

// 二叉树的最大深度
func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue, cur := make([]*TreeNode, 0), root
	queue = append(queue, cur)
	deep := 0
	for len(queue) != 0 {
		size := len(queue)
		deep++
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
	return deep
}

// 从中序前序遍历结果重建二叉树
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 {
		return nil
	}
	index := findIndex(inorder, preorder[0])

	return &TreeNode{
		Val:   preorder[0],
		Left:  buildTree(preorder[1:index+1], inorder[:index]),
		Right: buildTree(preorder[index+1:], inorder[index+1:]),
	}
}
func findIndex(nums []int, num int) int {
	for k, v := range nums {
		if v == num {
			return k
		}
	}
	return -1
}

// 复原 ip 地址
// dfs回溯
func restoreIpAddresses(s string) []string {
	var res, path []string
	backtrackRrestoreIP(s, path, 0, &res)
	return res
}
func backtrackRrestoreIP(s string, path []string, start int, res *[]string) {
	if start == len(s) && len(path) == 4 {
		tmp := strings.Join(path, ".")
		*res = append(*res, tmp)
	}
	for i := start; i < len(s); i++ {
		path = append(path, s[start:i+1])
		if i-start <= 4 && len(path) <= 4 && isValidIP(s, start, i) {
			backtrackRrestoreIP(s, path, i+1, res)
		} else {
			return
		}
		// 剪枝
		path = path[:len(path)-1]
	}
}
func isValidIP(s string, start int, end int) bool {
	check, _ := strconv.Atoi(s[start : end+1])
	// 前导为 0
	if end-start > 0 && s[start] == '0' {
		return false
	}
	if check > 255 {
		return false
	}
	return true
}
