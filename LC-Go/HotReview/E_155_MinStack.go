package HotReview

type MinStack struct {
	element, min []int
	len          int
}

/** initialize your data structure here. */
func Constructor155() MinStack {
	return MinStack{
		make([]int, 0),
		make([]int, 0),
		0,
	}
}

func (this *MinStack) Push(val int) {
	this.element = append(this.element, val)
	if this.len == 0 {
		this.min = append(this.min, val)
	} else {
		min := this.GetMin()
		if val < min {
			this.min = append(this.min, val)
		} else {
			this.min = append(this.min, min)
		}
	}
	this.len++
}

func (this *MinStack) Pop() {
	this.len--
	this.min = this.min[:this.len]
	this.element = this.element[:this.len]
}

func (this *MinStack) Top() int {
	return this.element[this.len-1]
}

func (this *MinStack) GetMin() int {
	return this.min[this.len-1]
}
