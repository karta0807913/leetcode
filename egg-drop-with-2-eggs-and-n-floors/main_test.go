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

/*
// original
func twoEggDrop(n int) int {
	cache := make([]int, 1001)
	cache[1] = 1
	cache[2] = 2

	for floor := 3; floor <= n; floor++ {
		min := math.MaxInt64
		for downfloor := 1; downfloor < floor; downfloor++ {
			tmp := max(downfloor, cache[floor-downfloor]+1)
			if tmp < min {
				min = tmp
			}
		}
		cache[floor] = min
	}
	return cache[n]
}
*/

func twoEggDrop(n int) int {
	cache := make([]int, n)
	cache[0] = 1

	for floor := 2; floor <= n; floor++ {
		min := math.MaxInt64
		for downfloor := 1; downfloor < floor; downfloor++ {
			left := downfloor
			right := cache[floor-downfloor-1] + 1
			tmp := max(left, right)
			// fmt.Printf("downfloor: %v, cache[floor-downfloor-1]+1: %v\n", downfloor, cache[floor-downfloor-1]+1)
			if tmp < min {
				min = tmp
			}
			if left > right {
				break
			}
		}
		cache[floor-1] = min
	}
	return cache[n-1]
}

func TestEggDrop(t *testing.T) {
	assert.Equal(t, 2, twoEggDrop(2))
	assert.Equal(t, 2, twoEggDrop(3))
	assert.Equal(t, 3, twoEggDrop(4))
	assert.Equal(t, 3, twoEggDrop(5))
	assert.Equal(t, 3, twoEggDrop(6))
	assert.Equal(t, 4, twoEggDrop(7))
	assert.Equal(t, 4, twoEggDrop(8))
	assert.Equal(t, 4, twoEggDrop(9))
	assert.Equal(t, 6, twoEggDrop(16))
	assert.Equal(t, 16, twoEggDrop(128))
	assert.Equal(t, 23, twoEggDrop(256))
}
