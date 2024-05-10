package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func sumByMask(cookies []int, mask int) int {
	index := 0
	sum := 0
	for ; mask != 0; mask >>= 1 {
		if (mask & 1) != 0 {
			sum += cookies[index]
		}
		index++
	}
	return sum
}

func distributeCookies(cookies []int, k int) int {
	totalSum := make([]int, 1<<len(cookies))
	dp := make([]int, 1<<len(cookies))
	next := make([]int, 1<<len(cookies))
	for i := 1; i < len(dp); i++ {
		dp[i] = sumByMask(cookies, i)
		totalSum[i] = dp[i]
	}
	// fmt.Printf("dp: %v\n", dp)
	k--
	for ; k != 0; k-- {
		for mask := 0; mask < len(dp); mask++ {
			total := totalSum[mask]
			ans := math.MaxInt
			for i := 0; i <= mask; i++ {
				ans = min(
					max(dp[i&mask], total-totalSum[i&mask]),
					ans,
				)
			}
			next[mask] = ans
		}
		dp, next = next, dp
		// fmt.Printf("dp: %v\n", dp)
	}
	return dp[len(dp)-1]
}

func TestCookies(t *testing.T) {
	assert := require.New(t)

	assert.Equal(
		7,
		distributeCookies([]int{
			6, 1, 3, 2, 2, 4, 1, 2,
		}, 3),
	)
	assert.Equal(
		38,
		distributeCookies([]int{
			15, 18, 19, 5, 6, 13, 15, 20,
		}, 3),
	)
	assert.Equal(
		31,
		distributeCookies([]int{8, 15, 10, 20, 8}, 2),
	)
}
