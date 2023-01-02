package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e9 + 7

func countPartitions(nums []int, k int) int {
	knapsack := make([]int, k)
	knapsack[0] = 1
	total := 0
	result := 1
	for _, num := range nums {
		for i := k - 1; i >= num; i-- {
			knapsack[i] += knapsack[i-num]
			knapsack[i] %= mod
		}
		total += num
		result *= 2
		result %= mod
	}
	if total < 2*k {
		return 0
	}
	sub := 0
	for _, val := range knapsack {
		sub += val * 2
		sub %= mod
	}
	return (result + mod - sub) % mod
}

func TestPartitions(t *testing.T) {
	assert := require.New(t)
	assert.Equal(0, countPartitions([]int{
		96, 40, 22, 98, 9, 97, 45, 22, 79, 57, 95, 62,
	}, 505))
}
