package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
)

func maxSlidingWindow(nums []int, k int) (result []int) {
	window := make([]int, 0)
	result = make([]int, 0, len(nums)+k)
	for i := range nums {
		window = append(window, i)
		end := len(window) - 2
		for ; end >= 0 && nums[window[end]] <= nums[window[end+1]]; end-- {
		}
		window[end+1] = window[len(window)-1]
		window = window[:end+2]
		fmt.Printf("window: %v\n", window)
		if window[0] <= i-k {
			window = window[1:]
		}
		result = append(result, nums[window[0]])
	}
	return result[k-1:]
}

func TestSlidingWindow(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		[]int{3, 3, 5, 5, 6, 7},
		maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3),
	)
	assert.Equal(
		[]int{3, 3, 5, 5, 6, 7},
		maxSlidingWindow([]int{1, 3, -1, -3, 5, 3, 6, 7}, 3),
	)
	assert.Equal(
		[]int{1},
		maxSlidingWindow([]int{1}, 1),
	)
	assert.Equal(
		[]int{1, 2, 3, 4, 5},
		maxSlidingWindow([]int{1, 2, 3, 4, 5}, 1),
	)

	assert.Equal(
		[]int{68, 24},
		maxSlidingWindow([]int{68, -81, 0, 24}, 3),
	)
}
