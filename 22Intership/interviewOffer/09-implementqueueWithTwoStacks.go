package interviewoffer

// 双栈模拟队列
type CQueue struct {
	in  []int
	out []int
}

func Constructor() CQueue {
	return CQueue{
		in:  []int{},
		out: []int{},
	}
}

func (this *CQueue) AppendTail(value int) {
	this.in = append(this.in, value)
}

func (this *CQueue) DeleteHead() int {
	if this.isEmpty() {
		return -1
	}
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

func (this *CQueue) isEmpty() bool {
	if len(this.in) == 0 && len(this.out) == 0 {
		return true
	}
	return false
}
