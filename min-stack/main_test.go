package main

type MinStack struct {
	MinStack []int
	Stack    []int
}

func Constructor() MinStack {
	return MinStack{}
}

func (this *MinStack) Push(val int) {
	mL := len(this.MinStack)
	if mL == 0 || this.Stack[this.MinStack[mL-1]] > val {
		this.MinStack = append(this.MinStack, len(this.Stack))
	}
	this.Stack = append(this.Stack, val)
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	if this.MinStack[len(this.MinStack)-1] != len(this.Stack) {
		return
	}
	this.MinStack = this.MinStack[:len(this.MinStack)-1]
}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) GetMin() int {
	return this.Stack[this.MinStack[len(this.MinStack)-1]]
}
