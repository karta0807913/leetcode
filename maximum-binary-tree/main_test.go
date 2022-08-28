package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func constractTree(nums []int, stack []int) *TreeNode {
	if len(nums) == 0 {
		return nil
	}
	i := 0
	for ; i+stack[i] < len(nums) && stack[i] != -1; i += stack[i] {
	}
	return &TreeNode{
		Val:   nums[i],
		Left:  constractTree(nums[:i], stack[:i]),
		Right: constractTree(nums[i+1:], stack[i+1:]),
	}
}

func constructMaximumBinaryTree(nums []int) *TreeNode {
	nextBigger := make([]int, len(nums))
	monoStack := make([]int, 0, len(nums))
	for i := len(nums) - 1; i >= 0; i-- {
		currentNum := nums[i]
		for len(monoStack) != 0 {
			lastItem := len(monoStack) - 1
			if nums[monoStack[lastItem]] > currentNum {
				break
			}
			monoStack = monoStack[:lastItem]
		}
		if len(monoStack) == 0 {
			nextBigger[i] = -1
		} else {
			nextBigger[i] = monoStack[len(monoStack)-1] - i
		}
		monoStack = append(monoStack, i)
		// fmt.Printf("monoStack: %v\n", monoStack)
	}
	// fmt.Println("HI")
	// fmt.Printf("nextBigger: %v\n", nextBigger)
	return constractTree(nums, nextBigger)
}

func TestBinaryTree(t *testing.T) {
	assert := require.New(t)
	assert.Equal(&TreeNode{
		Val: 6,
		Left: &TreeNode{
			Val: 3,
			Right: &TreeNode{
				Val: 2,
				Right: &TreeNode{
					Val: 1,
				},
			},
		},
		Right: &TreeNode{
			Val: 5,
			Left: &TreeNode{
				Val: 0,
			},
		},
	}, constructMaximumBinaryTree([]int{
		3, 2, 1, 6, 0, 5,
	}))
}
