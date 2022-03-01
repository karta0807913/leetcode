package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func candy(ratings []int) int {
	candys := make([]int, len(ratings))
	candy := 0
	prevRate := 0
	total := 0
	for i, rate := range ratings {
		if rate > prevRate {
			candy += 1
		} else {
			candy = 1
		}
		prevRate = rate
		candys[i] = candy
	}

	candy = 0
	prevRate = 0
	for idx := len(ratings) - 1; idx >= 0; idx-- {
		if ratings[idx] > prevRate {
			candy += 1
		} else {
			candy = 1
		}
		prevRate = ratings[idx]
		candys[idx] = max(candys[idx], candy)
		total += candys[idx]
	}

	return total
}

func TestCandy(t *testing.T) {
	assert.Equal(t, 4, candy([]int{1, 2, 2}))
	assert.Equal(t, 5, candy([]int{1, 0, 2}))
}
