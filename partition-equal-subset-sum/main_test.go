package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canPartition(nums []int) bool {
	sum := 0
	for _, num := range nums {
		sum += num
	}
	if sum%2 != 0 {
		return false
	}
	sum /= 2
	dp := make([]bool, sum+1)
	dp[0] = true
	for _, num := range nums {
		for i := sum; i >= num; i-- {
			if dp[i-num] {
				dp[i] = true
			}
		}
		if dp[sum] {
			return true
		}
	}
	return false
}

func TestPartitioin(t *testing.T) {
	assert.False(t, canPartition([]int{1, 2, 5}))
}
