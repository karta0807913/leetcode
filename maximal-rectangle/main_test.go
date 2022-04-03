package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximalRectangle(matrix [][]byte) int {
	area := make([]int, len(matrix[0]))
	height := make([]int, len(matrix[0]))
	stack := make([]int, 0, len(matrix[0])+1)
	maxHeight := 0
	for _, column := range matrix {
		for idx := range column {
			if column[idx] == '1' {
				height[idx]++
			} else {
				height[idx] = 0
			}
			area[idx] = 0
			for n := len(stack) - 1; n >= 0; n-- {
				i := stack[n]
				if height[i] < height[idx] {
					area[idx] = height[idx] * (idx - i)
					stack = append(stack, idx)
					maxHeight = max(maxHeight, area[idx])
					break
				} else if height[i] > height[idx] {
					area[i] += height[i] * (idx - i - 1)
					maxHeight = max(maxHeight, area[i])
				} else {
					area[idx] = height[idx] * (idx - i + 1)
					maxHeight = max(maxHeight, area[idx])
					// stack = stack[:n]
					stack = append(stack, idx)
					break
				}
				stack = stack[:n]
			}
			if len(stack) == 0 {
				stack = append(stack, idx)
				area[idx] = height[idx] * (idx + 1)
				maxHeight = max(maxHeight, area[idx])
			}
			// fmt.Printf("area: %v\n", area)
		}
		// fmt.Printf("stack: %v\n", stack)
		// fmt.Printf("area: %v\n", area)
		for _, i := range stack {
			area[i] += (len(area) - i - 1) * height[i]
			maxHeight = max(maxHeight, area[i])
		}
		stack = stack[:0]
		// fmt.Printf("area: %v\n", area)
	}
	return maxHeight
}

func TestRect(t *testing.T) {
	assert.Equal(t, 2, maximalRectangle([][]byte{
		{'1', '1'},
	}))
	assert.Equal(t, 6, maximalRectangle([][]byte{
		{'1', '0', '1', '1', '1'},
		{'0', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '0'},
		{'1', '1', '0', '1', '0'},
	}))
	assert.Equal(t, 4, maximalRectangle([][]byte{
		{'1', '0', '1', '1', '0'},
		{'0', '0', '1', '1', '1'},
		{'1', '0', '1', '0', '0'},
		{'1', '1', '0', '1', '0'},
	}))
	assert.Equal(t, 4, maximalRectangle([][]byte{
		{'1', '0', '1', '1', '0'},
		{'0', '1', '1', '1', '1'},
		{'1', '0', '1', '0', '0'},
		{'1', '1', '0', '1', '0'},
	}))
	assert.Equal(t, 6, maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
	assert.Equal(t, 10, maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '1', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}))
	assert.Equal(t, 1, maximalRectangle([][]byte{
		{'1', '0', '1', '0', '0'},
	}))
	assert.Equal(t, 0, maximalRectangle([][]byte{
		{'0', '0', '0'},
	}))
}
