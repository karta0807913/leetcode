package main

import (
	"fmt"
	"math"
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

type DP struct {
	Min   int
	Max   int
	Price int
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maximumTastiness(price []int, k int) int {
	sort.Ints(price)
	dp := make([]DP, k)
	dp[0] = DP{
		Min:   math.MaxInt,
		Max:   0,
		Price: price[0],
	}
	fmt.Printf("price: %v\n", price)
	for priceIndex := 1; priceIndex < len(price); priceIndex++ {
		for i := min(k-1, priceIndex); i >= 1; i-- {
			if n := min(price[priceIndex]-dp[i-1].Price, dp[i-1].Min); n > dp[i].Min {
				dp[i].Min = n
				dp[i].Price = price[priceIndex]
			}
		}
		fmt.Printf("dp: %v\n", dp)
	}
	return dp[k-1].Min
}

func TestTTT(t *testing.T) {
	assert := require.New(t)
	assert.Equal(19, maximumTastiness([]int{
		34, 116, 83, 15, 150, 56, 69, 42, 26,
	}, 6))
	assert.Equal(8, maximumTastiness([]int{
		13, 5, 1, 8, 21, 2,
	}, 3))
	assert.Equal(2, maximumTastiness([]int{
		1, 3, 1,
	}, 2))
	assert.Equal(0, maximumTastiness([]int{
		7, 7, 7, 7,
	}, 3))
}
