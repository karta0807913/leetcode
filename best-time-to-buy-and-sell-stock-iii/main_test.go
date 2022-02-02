package main

import (
	"math"
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
	sell1, sell2, buy1, buy2 := 0, 0, math.MinInt64, math.MinInt64
	for _, p := range prices {
		buy1 = max(buy1, -p)
		sell1 = max(sell1, buy1+p)
		buy2 = max(buy2, sell1-p)
		sell2 = max(sell2, buy2+p)
	}

	return sell2
}

func TestProfit(t *testing.T) {
	assert.Equal(t, 6+1, maxProfit([]int{
		3, 3, 5, 0, 0, 3, 1, 4,
	}))
}
