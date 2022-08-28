package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func longestMountain(arr []int) int {
	delta := 0
	length := 0
	maxLength := 0
	prev := arr[0]
	for i := 1; i < len(arr); i++ {
		height := arr[i]
		if prev < height {
			if delta == 1 {
				length += 1
			} else {
				maxLength = max(maxLength, length)
				length = 2
			}
			delta = 1
		} else if prev > height {
			if delta == 1 || delta == -1 {
				length += 1
				delta = -1
			}
		} else {
			if delta == -1 {
				maxLength = max(maxLength, length)
			}
			delta = 0
			length = 0
		}
		prev = height
		// fmt.Printf("length: %v, delta: %v, height: %v\n", length, delta, height)
	}
	if delta == -1 {
		maxLength = max(maxLength, length)
	}
	return maxLength
}

func TestMountain(t *testing.T) {
	assert := require.New(t)
	assert.Equal(3, longestMountain([]int{
		0, 0, 1, 0, 0, 1, 1, 1, 1, 1,
	}))
	assert.Equal(5, longestMountain([]int{
		2, 1, 4, 7, 3, 2, 5,
	}))
	assert.Equal(0, longestMountain([]int{
		2, 2, 2,
	}))
	assert.Equal(0, longestMountain([]int{
		2, 1, 0,
	}))
	assert.Equal(0, longestMountain([]int{
		1, 2, 3,
	}))
	assert.Equal(3, longestMountain([]int{
		1, 2, 1,
	}))
	assert.Equal(0, longestMountain([]int{
		1, 2, 2, 1,
	}))
}
