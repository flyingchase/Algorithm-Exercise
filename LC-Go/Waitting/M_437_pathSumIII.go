package Waitting
//给定一个二叉树的根节点 root ，和一个整数 targetSum ，求该二叉树里节点值之和等于 targetSum 的 路径 的数目。
//路径 不需要从根节点开始，也不需要在叶子节点结束，但是路径方向必须是向下的（只能从父节点到子节点）。
// #1 dfs+树的遍历算法
	var res int
func pathSum437(root *TreeNode, targetSum int) int {
	dfs437_1(root,targetSum)
	return res
}

func dfs437_1(root *TreeNode, sum int) {
	if root==nil {
		return
	}

	dfs437_2(root,root.Val,sum)
	dfs437_1(root.Left, sum)
	dfs437_1(root.Right, sum)
}

func dfs437_2(root *TreeNode, val int, sum int) {
	if val==sum{
		res++
	}
	if root.Left!=nil{
		dfs437_2(root.Left,val+root.Left.Val,sum)
	}
	if root.Right != nil {
		dfs437_2(root.Right,root.Right.Val+val,sum)
	}
}
// 前缀树
//
//func pathSum0(root *TreeNode, targetsum int) int  {
//	prefixsum:=make(map[int]int)
//	prefixsum[0]=1
//	return dfs437prefix(root,prefixsum,0,targetsum)
//
//}
