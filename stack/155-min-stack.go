package stack

type MinStack struct {
	min []int
	stack []int
	top int
}


func Constructor() MinStack {
	return MinStack{
		min: []int{},
		stack: []int{},
		top: -1,
	}
}

func (this *MinStack) Push(val int) {
	if this.top == -1 {
		this.min = append(this.min, val)
	} else {
		this.min = append(this.min, min(val, this.min[this.top]))
	}
	this.stack = append(this.stack, val)
	this.top++
}

func (this *MinStack) Pop() {
	this.stack = this.stack[:this.top]
	this.min = this.min[:this.top]
	this.top--
}

func (this *MinStack) Top() int {
	return this.stack[this.top]
}

func (this *MinStack) GetMin() int {
	return this.min[this.top]
}