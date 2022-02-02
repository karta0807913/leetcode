package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canFinish(piles []int, h, count int) bool {
	for _, num := range piles {
		n := num / count
		h -= n
		if n*count != num {
			h -= 1
		}
		if h < 0 {
			return false
		}
	}
	return true
}

func minEatingSpeed(piles []int, h int) int {
	right := 0
	for _, num := range piles {
		if num > right {
			right = num
		}
	}
	left := 1
	for left <= right {
		mid := (left + right) / 2
		if canFinish(piles, h, mid) {
			right = mid - 1
		} else {
			left = mid + 1
		}
	}
	return left
}

func TestMin(t *testing.T) {
	assert.Equal(t, 4, minEatingSpeed([]int{
		3, 6, 7, 11,
	}, 8))
}
