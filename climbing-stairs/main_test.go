package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func climbStairs(n int) int {
	j, k := 1, 1
	for i := 2; i <= n; i++ {
		j, k = k, j+k
	}
	return k
}

func TestFunc(t *testing.T) {
	assert.Equal(t, 2, climbStairs(2))
	assert.Equal(t, 3, climbStairs(3))
	assert.Equal(t, 5, climbStairs(4))
	assert.Equal(t, 8, climbStairs(5))
}
