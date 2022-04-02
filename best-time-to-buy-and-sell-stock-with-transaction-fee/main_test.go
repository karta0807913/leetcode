package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func maxProfit(prices []int, fee int) int {
	holdState := math.MinInt64
	selledState := 0
	for i := range prices {
		nextHold := selledState - prices[i] - fee
		nextSell := holdState + prices[i]
		if holdState < nextHold {
			holdState = nextHold
		}
		if selledState < nextSell {
			selledState = nextSell
		}
	}
	return selledState
}

func TestProfit(t *testing.T) {
	assert.Equal(t, 8, maxProfit([]int{1, 3, 2, 8, 4, 9}, 2))
	assert.Equal(t, 6, maxProfit([]int{1, 3, 7, 5, 10, 3}, 3))
}
