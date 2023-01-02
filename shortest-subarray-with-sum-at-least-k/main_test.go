package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

type Node struct {
	Idx int
	Val int
}

func shortestSubarray(nums []int, k int) int {
	stack := make([]Node, 0, len(nums))
	sum := 0
	result := math.MaxInt
	for i, num := range nums {
		sum += num
		if sum >= k && i < result {
			result = i + 1
		}
		for len(stack) != 0 && stack[len(stack)-1].Val >= sum {
			stack = stack[:len(stack)-1]
		}
		for len(stack) != 0 && sum-stack[0].Val >= k {
			sub := i - stack[0].Idx
			if result > sub {
				result = sub
			}
			stack = stack[1:]
		}
		stack = append(stack, Node{
			Idx: i,
			Val: sum,
		})
	}
	if result == math.MaxInt {
		return -1
	}
	return result
}

func TestArray(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		2,
		shortestSubarray([]int{
			17, 85, 93, -45, -21,
		}, 150),
	)
	assert.Equal(
		1,
		shortestSubarray([]int{1}, 1),
	)
}
