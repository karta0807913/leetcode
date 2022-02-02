package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countPrimes(n int) (ans int) {
	if n < 2 {
		return 0
	}
	dp := make([]bool, n)
	dp[0] = true
	dp[1] = true
	for number := 2; number < n; number++ {
		if !dp[number] {
			ans++
			for i := number * number; i < n; i += number {
				dp[i] = true
			}
		}
	}
	return
}

func TestCount(t *testing.T) {
	assert.Equal(t, 4, countPrimes(10))
	assert.Equal(t, 0, countPrimes(2))
}
