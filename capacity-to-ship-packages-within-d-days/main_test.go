package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canShip(weights []int, days int, capacity int) bool {
	c := 0
	for _, w := range weights {
		if capacity < w {
			return false
		}
		if c+w > capacity {
			days -= 1
			c = 0
			if days < 0 {
				return false
			}
		}
		c += w
	}
	days -= 1
	return days >= 0
}

func shipWithinDays(weights []int, days int) int {
	left, right := 1, 0
	for _, w := range weights {
		right += w
	}
	for left <= right {
		mid := (left + right) / 2
		if canShip(weights, days, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}
func TestCan(t *testing.T) {
	assert.False(t, canShip([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 1))
	assert.False(t, canShip([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 10))
	assert.False(t, canShip([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5, 11))
	assert.Equal(t, 15, shipWithinDays([]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5))
}
