package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func integerBreak(n int) int {
	cache := make([]int, n+1)
	cache[1] = 1
	cache[2] = 1
	for num := 3; num <= n; num++ {
		for leftNum := 1; leftNum < num; leftNum++ {
			// fmt.Printf("leftNum: %v, rightNum: %v\n", leftNum, num-leftNum)
			rightNum := num - leftNum
			tmp := leftNum * cache[rightNum]
			tmp = max(leftNum*rightNum, tmp)
			if cache[num] < tmp {
				cache[num] = tmp
			}
		}
	}
	fmt.Printf("cache: %v\n", cache)
	return cache[n]
}

func TestBreak(t *testing.T) {
	assert.Equal(t, 2, integerBreak(3))
	assert.Equal(t, 4, integerBreak(4))
	assert.Equal(t, 6, integerBreak(5))
	assert.Equal(t, 9, integerBreak(6))
	assert.Equal(t, 12, integerBreak(7))
	assert.Equal(t, 36, integerBreak(10))
}
