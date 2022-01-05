package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}
func maxProfit(prices []int) int {
	prev, current := 0, 0
	for i := 1; i < len(prices); i++ {
		delta := prices[i] - prices[i-1]
		prev, current = max(current, prev), max(current+delta, prev)
	}
	// fmt.Printf("current: %v, prev: %v\n", current, prev)
	return max(current, prev)
}

func TestProfit(t *testing.T) {
	assert.Equal(t, 3, maxProfit([]int{1, 2, 3, 0, 2}))
	assert.Equal(t, 0, maxProfit([]int{1}))
	assert.Equal(t, 5, maxProfit([]int{1, 2, 3, 4, 5, 6}))
	assert.Equal(t, 3, maxProfit([]int{1, 4, 2}))
	assert.Equal(t, 7, maxProfit([]int{6, 1, 6, 4, 3, 0, 2}))
}
