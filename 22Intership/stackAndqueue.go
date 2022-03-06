package intership

// ========================================================================
// 双栈实现队列
type MyQueue struct {
	in  []int
	out []int
}

func (queue *MyQueue) push(x int) {
	queue.in = append(queue.in, x)
}
func (queue *MyQueue) pop() int {
	// 确保有元素
	queue.peek()
	return queue.out[len(queue.out)-1]
}
func (queue *MyQueue) peek() int {
	if len(queue.out) == 0 {
		for len(queue.in) != 0 {
			queue.out = append(queue.out, queue.in[len(queue.in)-1])
			queue.in = queue.in[:len(queue.in)-1]
		}
	}
	return queue.out[len(queue.out)-1]
}
func (queue *MyQueue) isEmpty() bool {
	return len(queue.in) == 0 && len(queue.out) == 0
}

// ========================================================================
// 队列实现栈
type MyStack struct {
	queue []int
}

func (stack *MyStack) pop() int {
	node := stack.queue[0]
	stack.queue = stack.queue[1:]
	return node
}
func (stack *MyStack) push(x int) {
	tmp := append([]int{}, stack.queue...)
	stack.queue = []int{}
	stack.queue = append(stack.queue, x)
	stack.queue = append(stack.queue, tmp...)
}

// ========================================================================
// 有效括号序列

// ========================================================================
// 下一个更大的数

// minStack

type MinStack struct {
	data []int
	mini []int
}

func (m *MinStack) Push(x int) {
	m.data = append(m.data, x)
	if len(m.mini) == 0 || (len(m.mini) != 0 && m.mini[len(m.mini)-1] > x) {
		m.mini = append(m.mini, x)
	}
}
func (m *MinStack) Pop() int {
	if len(m.data) != 0 {
		if m.data[len(m.data)] == m.mini[len(m.mini)] {
			x := m.mini[len(m.mini)]
			m.mini = m.mini[:len(m.mini)]
			m.data = m.data[:len(m.data)]
			return x
		}
		x := m.mini[len(m.mini)-1]
		m.mini = m.mini[:len(m.mini)-1]
		return x
	}
	return -1
}

// 滑动窗口平均值
func MovingAverage(nums []int, size int) []float64 {
	if len(nums) < size {
		return nil
	}
	queue := make([]int, 0)
	sum := 0
	res := []float64{}
	for i := 0; i < len(nums); i++ {
		queue = append(queue, nums[i])
		sum += nums[i]
		if len(queue) > size {
			sum -= queue[0]
			queue = queue[1:]
		}
		if len(queue) == size {
			res = append(res, float64(sum)/float64(size))
		}
	}
	return res
}
