package days

import (
	"math"
)

// 验证对称二叉树
func isSymmetric(root *TreeNode) bool {
	if root == nil {
		return false
	}
	left, right := root.Left, root.Right
	return judge(left, right)
}
func judge(l, r *TreeNode) bool {
	if l == nil && r == nil {
		return true
	}
	if l == nil || r == nil {
		return false
	}
	if l.Val != r.Val {
		return false
	}
	return judge(r.Left, l.Right) && judge(r.Right, l.Left)
}

// min 栈
type MinStack struct {
	data []int
	min  []int
}

func Constructor() MinStack {
	return MinStack{
		data: []int{},
		min:  []int{math.MaxInt32},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	top := this.min[len(this.min)-1]
	this.min = append(this.min, min(top, val))
}

func (this *MinStack) Pop() {
	this.min = this.min[:len(this.min)-1]
	this.data = this.data[:len(this.data)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}

// 逆转链表，使用栈
func isPalindrome1(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	count := 0
	cur := head
	for cur != nil {
		count++
		cur = cur.Next
	}
	mid := count / 2
	cur = head
	stack := []*ListNode{}
	for mid > 0 {
		mid--
		stack = append(stack, cur)
		cur = cur.Next
	}
	if count%2 == 1 {
		cur = cur.Next
	}
	for cur != nil {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if node.Val != cur.Val {
			return false
		}
		cur = cur.Next
	}
	return true
}

// 单向遍历不断逆转
func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}
	fast, slow := head, head
	var prev *ListNode = nil
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		next := slow.Next
		slow.Next = prev
		prev, slow = slow, next
	}
	// odd
	if fast != nil {
		slow = slow.Next
	}
	for slow != nil {
		if slow.Val != prev.Val {
			return false
		}
		slow, prev = slow.Next, prev.Next
	}
	return true
}

// 验证搜索二叉树
func isValidBST(root *TreeNode) bool {
	if root == nil {
		return false
	}
	stack, cur := []*TreeNode{}, root
	res := []int{}
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
	for i := 0; i < len(res)-1; i++ {
		if res[i] > res[i+1] {
			return false
		}
	}
	return true
}

// 数学问题
/*
(randX()-1)*Y+randY() 可生成 1-x*y 之间的随机数
再通过取余操作即可获得 10 以内的随机数
*/
func rand10() int {
	for {
		x := rand7()
		y := rand7()
		z := (x-1)*7 + y
		if z <= 40 {
			return z%10 + 1
		}
		// 剩下 41-49 使用 rand9()=z-40
		a := rand7()
		b := z - 40
		c := (b-1)*7 + a
		if c <= 60 {
			return c%10 + 1
		}
		// 剩下 61-63 使用 rand3()=c-60
		e := rand7()
		f := c - 60
		g := (f-1)*7 + e
		if g <= 20 {
			return g%10 + 1
		}
		// 1的时候循环
	}

}
