package main

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// Stack solution
type BSTIterator struct {
	stack []*TreeNode
}

func Constructor(current *TreeNode) (iterator BSTIterator) {
	iterator.push(current)
	return
}

func (this *BSTIterator) Next() int {
	n := len(this.stack) - 1
	pop := this.stack[n]
	this.push(pop.Right)
	this.stack = this.stack[:n]
	return pop.Val
}

func (this *BSTIterator) HasNext() bool {
	return len(this.stack) != 0
}

func (this *BSTIterator) push(node *TreeNode) {
	for node != nil {
		this.stack = append(this.stack, node)
		node = node.Left
	}
}

/*
// Morris traversal
type LinkList struct {
	Val  int
	Next *LinkList
}

type BSTIterator struct {
	current *LinkList
}

func Constructor(current *TreeNode) (iterator BSTIterator) {
	tail := &iterator.current
	for current != nil {
		if current.Left != nil {
			runner := &current.Left
			for *runner != nil && *runner != current {
				runner = &(*runner).Right
			}
			if *runner == nil {
				*runner = current
				current = current.Left
			} else {
				*tail = &LinkList{
					Val: current.Val,
				}
				tail = &(*tail).Next
				current = current.Right
			}
		} else {
			*tail = &LinkList{
				Val: current.Val,
			}
			tail = &(*tail).Next
			current = current.Right
		}
	}
	return
}

func (this *BSTIterator) Next() int {
	val := this.current.Val
	this.current = this.current.Next
	return val
}

func (this *BSTIterator) HasNext() bool {
	return this.current != nil
}
*/
