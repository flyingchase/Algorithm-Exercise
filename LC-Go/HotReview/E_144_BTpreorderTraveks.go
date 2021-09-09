package HotReview

import (
	"fmt"
)

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func preorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	cur:=root

	res:=make([]int,0)
	stack:=make([]*TreeNode,0)
	for cur!=nil||len(stack)>0 {
		for cur!=nil {
			// 前序遍历即先将 cur 存储起来
			res=append(res,cur.Val)
			stack=append(stack,cur)
			// 每次往最左侧走
			cur=cur.Left
		}
		node:=stack[len(stack) - 1]
		stack=stack[:len(stack) - 1]
		// 左侧走完后遍历右侧
		cur=node.Right
	}
	return res
}

func postOrderTraversal(root *TreeNode)[]int  {
	if root == nil {
		return nil
	}

	cur:=root
	res:=make([]int,0)
	stack:=make([]*TreeNode,0)

	for cur != nil||len(stack)>0 {
		for cur != nil {
			res= append(res,cur.Val)
			stack=append(stack,cur)
			cur=cur.Right
		}
		node:=stack[len(stack) - 1]
		stack=stack[:len(stack) - 1]
		cur=node.Left
	}

	resTrue:=make([] int,len(res))
	for index:=range res {
	    resTrue[index]=res[len(res) - 1-index]
	}

	//for i:=0;i<len(res)>>1; i++{
	//    res[i],res[len(res) - 1-i]=res[len(res) - 1-i],res[i]
	//}

	fmt.Println("before: ",res)
	//sort.Reverse(sort.IntSlice{res})
	//sort.Ints(res)
	fmt.Println(res)
	return resTrue

}


