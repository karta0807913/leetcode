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

func largestRectangleArea(heights []int) int {
	area := make([]int, len(heights))
	stack := make([]int, 1, len(heights))
	stack[0] = 0
	area[0] = heights[0]
	maxArea := 0
	for i := 1; i < len(heights); i++ {
		h := heights[i]
		area[i] = h
		for n := len(stack) - 1; n >= 0; n-- {
			target := heights[stack[n]]
			if target < h {
				area[i] = h * (i - stack[n])
				break
			} else if target > h {
				area[stack[n]] += target * (i - stack[n] - 1)
				maxArea = max(maxArea, area[stack[n]])
			} else {
				break
			}
			stack = stack[:n]
		}
		if len(stack) == 0 {
			area[i] = h * (i + 1)
		}
		stack = append(stack, i)
		// fmt.Printf("stack: %v\n", stack)
		// fmt.Printf("area: %v\n", area)
	}
	for _, i := range stack {
		area[i] += heights[i] * (len(heights) - i - 1)
		maxArea = max(maxArea, area[i])
	}
	// fmt.Printf("area: %v\n", area)
	return maxArea
}

/*
// best?
func largestRectangleArea(heights []int) int {
	leftGreater := make([]int, len(heights))
	rightGreter := make([]int, len(heights))
	leftGreater[0] = -1
	for i := 1; i < len(leftGreater); i++ {
		p := i - 1
		for p >= 0 && heights[p] >= heights[i] {
			p = leftGreater[p]
		}
		leftGreater[i] = p
	}

	for i := len(rightGreter) - 1; i >= 0; i-- {
		p := i + 1
		for p < len(rightGreter) && heights[p] >= heights[i] {
			p = rightGreter[p]
		}
		rightGreter[i] = p
	}

	maxRect := 0
	for idx := range heights {
		rect := heights[idx] * (rightGreter[idx] - leftGreater[idx] - 1)
		if rect > maxRect {
			maxRect = rect
		}
	}
	return maxRect
}

/*
type Pair struct {
	Height, Idx int
}

// normal stack, faster
func largestRectangleArea(heights []int) int {
	monoStack := make([]Pair, 0, len(heights))
	maxRect := 0
	for idxR, heightR := range heights {
		minIdx := idxR
		if heightR > maxRect {
			maxRect = heightR
		}
		for idx, left := range monoStack {
			if left.Height < heightR {
				rect := left.Height * (idxR - left.Idx + 1)
				if rect > maxRect {
					maxRect = rect
				}
			} else {
				rect := heightR * (idxR - left.Idx + 1)
				minIdx = min(minIdx, left.Idx)
				monoStack = monoStack[:idx]
				if rect > maxRect {
					maxRect = rect
				}
				break
			}
		}
		monoStack = append(monoStack, Pair{Height: heightR, Idx: minIdx})
	}
	return maxRect
}
*/

func TestArea(t *testing.T) {
	// assert.Equal(t, 10, largestRectangleArea([]int{
	// 	2, 2, 2, 3, 2,
	// }))
	// assert.Equal(t, 6, largestRectangleArea([]int{
	// 	2, 3, 2,
	// }))
	assert.Equal(t, 10, largestRectangleArea([]int{
		2, 1, 5, 6, 2, 3,
	}))
}

/**
// map
func largestRectangleArea(heights []int) int {
	monoStack := make(map[int]int)
	maxRect := 0
	for idxR, right := range heights {
		minIdx := idxR
		if maxRect < right {
			maxRect = right
		}
		for left, idxL := range monoStack {
			rect := min(left, right) * (idxR - idxL + 1)
			if maxRect < rect {
				maxRect = rect
			}
			if left >= right {
				delete(monoStack, left)
				minIdx = min(minIdx, idxL)
			}
		}
		monoStack[right] = minIdx
	}
	return maxRect
}
**/
