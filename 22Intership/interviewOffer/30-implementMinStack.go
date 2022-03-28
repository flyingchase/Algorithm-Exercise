package interviewoffer

type MinStack struct {
	mini []int
	data []int
}

/** initialize your data structure here. */
func constructor() MinStack {
	return MinStack{mini: []int{}, data: []int{}}
}

func (this *MinStack) Push(x int) {
	this.data = append(this.data, x)
	if len(this.mini) == 0 || this.mini[len(this.mini)-1] > x {
		this.mini = append(this.mini, x)
	}
}

func (this *MinStack) Pop() {
	if len(this.mini) != 0 && this.mini[len(this.mini)-1] == this.data[len(this.data)-1] {
		this.data = this.data[:len(this.data)-1]
	}
	this.mini = this.mini[:len(this.mini)-1]
}

func (this *MinStack) Top() int {
	return this.data[len(this.data)-1]
}

func (this *MinStack) Min() int {
	return this.mini[len(this.mini)-1]
}
