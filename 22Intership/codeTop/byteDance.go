package codetop

import (
	"math"
)

func LengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	count := make(map[byte]int, 0)
	res := 0
	l, r := 0, 0
	for r < len(s) {
		if _, ok := count[s[r]]; !ok {
			count[s[r]]++
			r++
		} else {
			delete(count, s[l])
			l++
		}
		if res < r-l+1 {
			res = r - l + 1

		}
	}
	return res
}

// 最小覆盖子串
func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}
	dict := map[byte]int{}
	l, r, start, length := 0, 0, 0, 0
	length = math.MaxInt32
	match := 0
	for i := 0; i < len(t); i++ {
		dict[t[i]]++
	}
	for r < len(s) {
		if _, ok := dict[s[r]]; ok {
			dict[s[r]]--
			if dict[s[r]] == 0 {
				match++
			}
		}
		r++
		for match == len(dict) {
			if r-l < length {
				length = r - l
				start = l
			}
			if count, ok := dict[s[l]]; ok {
				if count == 0 {
					match--
				}
				// 划过了之后 dict 被删除的要加上
				dict[s[l]]++
			}
			l++
		}
	}
	if length == math.MaxInt32 {
		return ""
	}
	return s[start : start+length]
}

// 二叉树中路径和等于 target 的所有路径
func pathSum(root *TreeNode, targetSum int) [][]int {
	if root == nil {
		return nil
	}
	res := [][]int{}
	dfsPathSum(&res, root, targetSum, 0, []int{})
	return res
}
func dfsPathSum(res *[][]int, node *TreeNode, targetSum, curSum int, curPath []int) {
	curPath = append(curPath, node.Val)
	curSum += node.Val
	if curSum == targetSum && node.Left == nil && node.Right == nil {
		*res = append(*res, append([]int{}, curPath...))
		return
	}
	if node.Left != nil {
		dfsPathSum(res, node.Left, targetSum, curSum, curPath)
	}
	if node.Right != nil {
		dfsPathSum(res, node.Right, targetSum, curSum, curPath)
	}
	curPath = curPath[:len(curPath)-1]
}

// 编辑距离
func minDistance(word1 string, word2 string) int {
}

// 中序遍历
func InorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	stack, cur := make([]*TreeNode, 0), root
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
	return res
}

// 双栈实现队列
type MyQueue struct {
	inStack  []int
	outStack []int
}

func Constructor() MyQueue {
	return MyQueue{
		inStack:  []int{},
		outStack: []int{},
	}
}

func (this *MyQueue) Push(x int) {
	this.inStack = append(this.inStack, x)
}

func (this *MyQueue) Pop() int {
	if len(this.outStack) == 0 {
		for len(this.inStack) != 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	node := this.outStack[len(this.outStack)-1]
	this.outStack = this.outStack[:len(this.outStack)-1]
	return node
}

func (this *MyQueue) Peek() int {
	if len(this.outStack) == 0 {
		for len(this.inStack) != 0 {
			this.outStack = append(this.outStack, this.inStack[len(this.inStack)-1])
			this.inStack = this.inStack[:len(this.inStack)-1]
		}
	}
	node := this.outStack[len(this.outStack)-1]
	return node
}

func (this *MyQueue) Empty() bool {
	return len(this.outStack) == 0 && len(this.inStack) == 0
}

// 寻找峰值
func findPeakElement(nums []int) int {
	if len(nums) == 1 {
		return 0
	}
	leftPeek := map[int]struct{}{}
	min := math.MinInt32
	tmp := append([]int{}, min)
	nums = append(tmp, nums...)
	nums = append(nums, min)
	leftOrder, rightOrder := 1, len(nums)-2

	for leftOrder < len(nums) {
		if nums[leftOrder] > nums[leftOrder-1] {
			leftPeek[leftOrder] = struct{}{}
		}
		leftOrder++
	}
	for rightOrder >= 0 {
		if nums[rightOrder] > nums[rightOrder+1] {
			if _, exist := leftPeek[rightOrder]; exist {
				return rightOrder - 1
			}
		}
		rightOrder--
	}
	return -1
}

// 实现最小栈
type MinStack struct {
	min  []int
	data []int
}

func Constructor() MinStack {
	return MinStack{
		min:  []int{math.MaxInt32},
		data: []int{},
	}
}

func (this *MinStack) Push(val int) {
	this.data = append(this.data, val)
	if this.min[len(this.min)-1] > val {
		this.min = append(this.min, val)
	}
}

func (this *MinStack) Pop() {
	this.data = this.data[:len(this.data)-1]
	this.min = this.min[:len(this.min)-1]

}

func (this *MinStack) Top() int {
	node := this.data[len(this.data)-1]
	return node
}

func (this *MinStack) GetMin() int {
	return this.min[len(this.min)-1]
}
