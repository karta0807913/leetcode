package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

/*
Binary Tree Solution
Since the worse case is O(n^2*log(n)), this is not a good solution
*/
// type SequenceNode struct {
// 	Left         *SequenceNode
// 	Right        *SequenceNode
// 	CurrentValue int
// 	Weight       int
// }

// func Insert(parent **SequenceNode, node *SequenceNode) {
// 	if node == nil {
// 		return
// 	}
// 	if (*parent) == nil {
// 		(*parent) = node
// 		return
// 	}
// 	if node.CurrentValue < (*parent).CurrentValue {
// 		right := node.Right
// 		node.Right = nil
// 		Insert(&(*parent).Left, node)
// 		Insert(parent, right)
// 	} else if node.CurrentValue > (*parent).CurrentValue {
// 		left := node.Left
// 		node.Left = nil
// 		Insert(&(*parent).Right, node)
// 		Insert(parent, left)
// 	} else {
// 		(*parent).Weight += node.Weight
// 		Insert(parent, node.Left)
// 		Insert(parent, node.Right)
// 	}
// }

// func (node *SequenceNode) Count() int {
// 	if node == nil {
// 		return 0
// 	}
// 	return node.Left.Count() + node.Right.Count() + node.Weight

// }

// func countSmaller(nums []int) []int {
// 	ans := make([]int, len(nums))
// 	head := &SequenceNode{
// 		CurrentValue: nums[len(nums)-1],
// 		Weight:       1,
// 	}
// 	for i := len(nums) - 2; i >= 0; i-- {
// 		newHead := &SequenceNode{
// 			CurrentValue: nums[i],
// 			Weight:       1,
// 		}
// 		Insert(&newHead, head)
// 		ans[i] = newHead.Left.Count()
// 		head = newHead
// 	}
// 	head.Print()
// 	fmt.Println()
// 	return ans
// }

// func (node *SequenceNode) Print() {
// 	if node.Left != nil {
// 		node.Left.Print()
// 	}
// 	fmt.Printf("%v ", node.CurrentValue)
// 	if node.Right != nil {
// 		node.Right.Print()
// 	}
// }

type Item struct {
	Value    int
	Position *int
}

func mergeSort(item []*Item) []*Item {
	if len(item) <= 1 {
		return item
	}
	mid := len(item) / 2
	result := make([]*Item, 0, len(item))
	left := mergeSort(item[:mid])
	right := mergeSort(item[mid:])
	c := 0
	for len(left) != 0 && len(right) != 0 {
		if left[0].Value > right[0].Value {
			c += 1
			result = append(result, right[0])
			right = right[1:]
		} else {
			*left[0].Position += c
			result = append(result, left[0])
			left = left[1:]
		}
	}
	for _, item := range left {
		*item.Position += c
		result = append(result, item)
	}
	result = append(result, right...)
	return result
}

func countSmaller(nums []int) []int {
	ans := make([]int, len(nums))
	data := make([]*Item, len(nums))
	for i := range nums {
		data[i] = &Item{
			Value:    nums[i],
			Position: &ans[i],
		}
	}
	mergeSort(data)
	return ans
}

func TestSmaller(t *testing.T) {
	assert := require.New(t)
	assert.Equal([]int{2, 1, 0}, countSmaller([]int{2, 1, 0}))
	assert.Equal([]int{2, 1, 1, 0}, countSmaller([]int{5, 2, 6, 1}))
	assert.Equal([]int{3, 1, 1, 0}, countSmaller([]int{2, 1, 1, 0}))
	assert.Equal([]int{0, 0}, countSmaller([]int{-1, -1}))
	assert.Equal([]int{2, 1, 1, 0}, countSmaller([]int{5, 2, 6, 1}))
	assert.Equal([]int{
		1, 1, 1, 0,
	}, countSmaller([]int{3, 9, 9, 1}))
	assert.Equal([]int{
		0, 2, 4, 2, 3, 0, 1, 1, 1, 0,
	}, countSmaller([]int{0, 2, 6, 2, 9, 1, 3, 9, 9, 1}))
}
