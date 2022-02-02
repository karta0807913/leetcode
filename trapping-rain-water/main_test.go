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
func trap(height []int) int {
	ans := make([]int, len(height))
	t := 0
	for i, h := range height {
		if h > t {
			t = h
		} else {
			ans[i] = t - h
		}
	}

	t = 0
	total := 0
	for i := len(height) - 1; i >= 0; i-- {
		if height[i] > t {
			t = height[i]
		} else {
			ans[i] = min(ans[i], t-height[i])
			total += ans[i]
		}
	}
	return total
}

func TestTrap(t *testing.T) {
	assert.Equal(t, 6, trap([]int{
		0, 1, 0, 2, 1, 0, 1, 3, 2, 1, 2, 1,
	}))
}
