package Waitting

import (
	"fmt"
	"strconv"
)

/*
关联 257 根节点到叶子结点的路径
129: 根节点到叶子结点所形成的的数字
*/
func sumNumber(root *TreeNode) int {
	res := 0
	// the parameter of dfs  res is pointer due to its type
	dfsSumNumber(root, 0, &res)
	return res
}

func dfsSumNumber(root *TreeNode, sum int, res *int) {
	if root == nil {
		return
	}
	sum = sum*10 + root.Val
	if root.Left == nil && root.Right == nil {
		*res += sum
		return
	}
	// due to  func has judge root==nil so no need to judge root.left/root.right ==nil or not
	dfsSumNumber(root.Left, sum, res)
	dfsSumNumber(root.Right, sum, res)
}

// 257 显示出根节点到叶子结点的路径 用“——>”链接
func binaryTreePaths(root *TreeNode) []string {
	res := make([]string, 0)
	// res 为 string 的数组 传入指针
	dfsbinaryTreePaths(root, &res, "")
	return res
}

func dfsbinaryTreePaths(root *TreeNode, res *[]string, temp string) {
	if root == nil {
		return
	}
	// judge if the temp connect the first str 首处理
	if temp == "" {
		temp += string(strconv.Itoa(root.Val))
	} else {
		temp += ("->" + string(strconv.Itoa(root.Val)))
	}
	if root.Left == nil && root.Right == nil {
		*res = append(*res, temp)
	}
	dfsbinaryTreePaths(root.Left, res, temp)
	dfsbinaryTreePaths(root.Right, res, temp)
}

// 257 使用 chan 无缓存
func binaryTreePaths2(root *TreeNode) []string {
	ch := make(chan string)
	go func() {
		defer close(ch)
		dfsBTPath2(root, "", ch)
	}()
	res := make([]string, 0)
	for c := range ch {
		res = append(res, c)
	}
	return res
}

func dfsBTPath2(root *TreeNode, s string, ch chan string) {
	if root == nil {
		return
	}
	if root.Left == root.Right && root.Right == nil {
		// 最后一份无须加 ->
		ch <- s + strconv.Itoa(root.Val)
	}
	s += fmt.Sprintf("%d->", root.Val)
	dfsBTPath2(root.Left, s, ch)
	dfsBTPath2(root.Right, s, ch)
}
