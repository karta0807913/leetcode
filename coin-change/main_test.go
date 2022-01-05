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

func coinChange(coins []int, amount int) int {
	if amount == 0 {
		return 0
	}
	cache := make([]int, amount+1)
	cache[0] = 0
	for _, coin := range coins {
		if coin >= len(cache) {
			continue
		}
		cache[coin] = 1
		// fmt.Printf("cache: %v\n", cache)
		for idx := coin + 1; idx < len(cache); idx++ {
			if cache[idx-coin] == 0 {
			} else if cache[idx] == 0 {
				cache[idx] = cache[idx-coin] + 1
			} else {
				cache[idx] = min(cache[idx], cache[idx-coin]+1)
			}

		}
		// fmt.Printf("cache: %v\n", cache)
	}
	if cache[amount] == 0 {
		return -1
	}
	return cache[amount]
}

func TestChange(t *testing.T) {
	assert.Equal(t, 3, coinChange([]int{1, 2, 5}, 11))
	assert.Equal(t, -1, coinChange([]int{2}, 3))
	assert.Equal(t, 1, coinChange([]int{2}, 2))
	assert.Equal(t, 5, coinChange([]int{2}, 10))
	assert.Equal(t, 0, coinChange([]int{2}, 0))
	assert.Equal(t, 0, coinChange([]int{2, 3, 4, 5, 6}, 0))
}
