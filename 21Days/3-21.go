package days

// 二叉树右视图
// queue 宽度优先遍历 保存每层最后一个
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res, queue, cur := []int{}, make([]*TreeNode, 0), root
	queue = append(queue, cur)
	for len(queue) != 0 {
		size := len(queue)
		tmp := []int{}
		for size > 0 {
			size--
			node := queue[0]
			queue = queue[1:]
			tmp = append(tmp, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		if len(tmp) != 0 {
			res = append(res, tmp[len(tmp)-1])
		}
	}
	return res
}

// 重排链表
// 后一半入栈，前一半入队，队列先弹出再谈栈顶，奇数时中间 mid 拼接到最后
func ReorderList(head *ListNode) {
	if head == nil {
		return
	}
	cur := head
	count := 0
	for cur != nil {
		count++
		cur = cur.Next
	}
	mid := count / 2
	queue, stack := []*ListNode{}, []*ListNode{}
	cur = head
	for i := 0; i < mid; i++ {
		queue = append(queue, cur)
		cur = cur.Next
	}
	next := cur.Next
	cur.Next = nil
	var last *ListNode = nil
	if count%2 == 1 {
		last = cur
	}
	for next != nil {
		stack = append(stack, next)
		next = next.Next
	}
	dummyHead := &ListNode{Val: -1}
	cur = dummyHead
	for len(queue) != 0 || len(stack) != 0 {
		if len(queue) != 0 {
			node := queue[0]
			queue = queue[1:]
			cur.Next = node
			cur = cur.Next
		}
		if len(stack) != 0 {
			node := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			cur.Next = node
			cur = cur.Next
		}
	}
	if last != nil {
		cur.Next = last
	}
	head = dummyHead.Next
}

func inorderTraversal(root *TreeNode) []int {
	if root == nil {
		return nil
	}
	res, cur, stack := []int{}, root, []*TreeNode{}
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

// 二分板子
func searchTemple(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	l, r := 0, len(nums)-1
	for l <= r {
		mid := l + (r-l)>>1
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			l = mid + 1
		} else {
			r = mid - 1
		}
	}
	return -1
}

// 双栈模拟队列
type MyQueue struct {
	in  []int
	out []int
}

func Constructor() MyQueue {
	return MyQueue{
		in:  []int{},
		out: []int{},
	}

}

func (this *MyQueue) Push(x int) {
	this.in = append(this.in, x)
}

func (this *MyQueue) Pop() int {
	if len(this.out) == 0 {
		for len(this.in) != 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
	node := this.out[len(this.out)-1]
	this.out = this.out[:len(this.out)-1]
	return node
}

func (this *MyQueue) Peek() int {
	if len(this.out) == 0 {
		for len(this.in) != 0 {
			this.out = append(this.out, this.in[len(this.in)-1])
			this.in = this.in[:len(this.in)-1]
		}
	}
	return this.out[len(this.out)-1]
}

func (this *MyQueue) Empty() bool {
	return len(this.in) == 0 && len(this.out) == 0
}
