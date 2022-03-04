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
	return 0
}
func (queue *MyQueue) peek() bool {
	return false
}

// ========================================================================
// 队列实现栈

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
		x := m.data[len(m.data)]
		m.data = m.data[:len(m.data)]
		return x
	}
	return -1
}
