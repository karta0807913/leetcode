package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func maxValueOfCoins(piles [][]int, k int) int {
	dp := make([]int, k+1)
	next := make([]int, k+1)
	for _, pile := range piles {
		// fmt.Printf("pile: %v\n", pile)
		sum := 0
		for i, coin := range pile {
			sum += coin
			for j := 0; j+i+1 <= k; j++ {
				next[j+i+1] = max(max(dp[j+i+1], sum+dp[j]), next[j+i+1])
			}
			// fmt.Printf("next: %v, sum: %d\n", next, sum)
		}
		dp, next = next, dp
		// fmt.Printf("dp: %v\n", dp)
	}
	return dp[k]
}

func TestCoins(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		494,
		maxValueOfCoins(
			[][]int{
				{37, 88},
				{51, 64, 65, 20, 95, 30, 26},
				{9, 62, 20},
				{44},
			},
			9,
		),
	)
	assert.Equal(
		101,
		maxValueOfCoins(
			[][]int{
				{1, 100, 3},
				{7, 8, 9},
			},
			2,
		),
	)
}
