package days

import (
	"math"
	"runtime/internal/math"
)

// 删除排序链表中的重复项，只保留非重复部分
func deleteDuplicates(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	dummyHead := &ListNode{Val: -1, Next: head}
	prev, cur := dummyHead, head
	for cur != nil {
		next := cur.Next
		flag := false
		for next != nil && next.Val == cur.Val {
			next = next.Next
			flag = true
		}
		if flag {
			prev.Next = next
			cur = next.Next
			continue
		}
		prev, cur = prev.Next, cur.Next
	}
	return dummyHead.Next
}

// 字符串转为整数
func myAtoi(s string) int {
	if len(s) == 0 {
		return 0
	}
	var index int
	size := len(s)
	for ; index < size && s[index] == ' '; index++ {
	}
	sign := 1
	res := 0
	if index < len(s) && s[index] == '-' {
		sign = -1
		index++
	} else if index < len(s) && s[index] == '+' {
		sign = 1
		index++
	}
	for ; index < len(s) && s[index] <= '9' && s[index] >= '0'; index++ {
		res = res*10 + int(s[index]-'0')
		if res*sign > math.MaxInt32 {
			return math.MaxInt32
		}
		if res*sign < math.MinInt32 {
			return math.MinInt32
		}
	}
	return res * sign
}

// 二叉树中的最大路径和
func maxPathSum(root *TreeNode) int {
	if root == nil {
		return 0
	}
	max := math.MinInt32
	// 最大和为l+root,r+root,l+r+root,root四个中的最大
	helperMaxPathSum(root, &max)
	return max
}
func helperMaxPathSum(node *TreeNode, max *int) int {
	if node == nil {
		return 0
	}
	l, r := maxMax(0, helperMaxPathSum(node.Left, max)), maxMax(0, helperMaxPathSum(node.Right, max))
	// 横向比较，停在当前节点
	*max = maxMax(*max, l+r+node.Val)
	// 纵向比较，左右子节点的最大值
	return maxMax(l, r) + node.Val
}
func maxMax(i, j int) int {
	if i >= j {
		return i
	}
	return j
}
