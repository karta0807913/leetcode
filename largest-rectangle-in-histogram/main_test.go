package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

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
