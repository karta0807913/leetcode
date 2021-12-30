package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit(prices []int) int {
	best := 0
	for idx := 1; idx < len(prices); idx++ {
		profit := prices[idx] - prices[idx-1]
		fmt.Printf("profit: %v\n", prices[idx]-prices[idx-1])
		if profit < 0 {
			profit = 0
		}
		best += profit
	}
	return best
}

func TestProfit(t *testing.T) {
	assert.Equal(t, 7, maxProfit([]int{7, 1, 5, 3, 6, 4}))
	assert.Equal(t, 4, maxProfit([]int{1, 2, 3, 4, 5}))
	assert.Equal(t, 0, maxProfit([]int{7, 6, 4, 3, 1}))
}
