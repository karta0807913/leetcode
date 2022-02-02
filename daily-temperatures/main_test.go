package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func dailyTemperatures(temperatures []int) (ans []int) {
	stack := make([]int, 0, len(temperatures)/2)
	ans = make([]int, len(temperatures))
	for idx := len(temperatures) - 1; idx >= 0; idx-- {
		for len(stack) != 0 {
			n := len(stack) - 1
			p := stack[n]
			stack = stack[:n]
			if temperatures[p] > temperatures[idx] {
				ans[idx] = p - idx
				stack = append(stack, p)
				break
			}
		}
		stack = append(stack, idx)
	}
	return
}

func TestTemp(t *testing.T) {
	assert.Equal(t, []int{
		1, 1, 4, 2, 1, 1, 0, 0,
	}, dailyTemperatures([]int{73, 74, 75, 71, 69, 72, 76, 73}))
}
