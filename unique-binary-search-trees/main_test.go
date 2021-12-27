package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func numTrees(n int) int {
	cache := make([]int, n+1)
	cache[0] = 1
	cache[1] = 1
	for i := 2; i <= n; i++ {
		for n := 0; n < i; n++ {
			cache[i] += cache[n] * cache[i-n-1]
		}
	}
	return cache[n]
}

func TestNumTrees(t *testing.T) {
	assert.Equal(t, 1, numTrees(1))
	assert.Equal(t, 2, numTrees(2))
	assert.Equal(t, 5, numTrees(3))
	assert.Equal(t, 14, numTrees(4))
	assert.Equal(t, 42, numTrees(5))
}
