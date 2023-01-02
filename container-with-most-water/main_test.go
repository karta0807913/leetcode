package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxArea(height []int) int {
	left, right := 0, len(height)-1
	maxArea := 0
	for left < right {
		maxArea = max((right-left)*min(height[left], height[right]), maxArea)
		if height[left] < height[right] {
			left++
		} else {
			right--
		}
	}
	return maxArea
}

func TestArea(t *testing.T) {
	assert := require.New(t)
	assert.Equal(49, maxArea(
		[]int{
			1, 8, 6, 2, 5, 4, 8, 3, 7,
		},
	))
}
