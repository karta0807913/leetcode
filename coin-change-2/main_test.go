package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func change(amount int, coins []int) int {
	var cache = make([]int, 5001)
	cache[0] = 1
	for _, coin := range coins {
		for idx := coin; idx <= amount; idx++ {
			cache[idx] += cache[idx-coin]
		}
	}
	return cache[amount]
}

func TestChange(t *testing.T) {
	assert.Equal(t, 0, change(3, []int{2}))
	assert.Equal(t, 1, change(1, []int{1, 2, 5}))
	assert.Equal(t, 2, change(2, []int{1, 2, 5}))
	assert.Equal(t, 2, change(3, []int{1, 2, 5}))
	assert.Equal(t, 3, change(4, []int{1, 2, 5}))
	assert.Equal(t, 4, change(5, []int{1, 2, 5}))
	assert.Equal(t, 5, change(6, []int{1, 2, 5}))
	assert.Equal(t, 717, change(500, []int{2, 7, 13}))
	assert.Equal(t, 1, change(9, []int{2, 7}))
	assert.Equal(t, 0, change(5000, []int{4001, 4002, 4003, 4004, 4005}))
}
