package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

// func max(a, b int) int {
// 	if a > b {
// 		return a
// 	}
// 	return b
// }

// func findLongestChain(pairs [][]int) int {
// 	sort.Slice(pairs, func(i, j int) bool {
// 		return pairs[i][1] < pairs[j][1]
// 	})
// 	// fmt.Printf("pairs: %v\n", pairs)
// 	dp := make([]int, len(pairs))
// 	prev := -2000
// 	for i := 0; i < len(pairs); i++ {
// 		if pairs[i][0] == prev {
// 			continue
// 		}
// 		prev = pairs[i][0]
// 		for j := i + 1; j < len(pairs); j++ {
// 			if pairs[i][1] < pairs[j][0] {
// 				dp[j] = max(dp[i]+1, dp[j])
// 			}
// 		}
// 		// fmt.Printf("dp: %v\n", dp)
// 	}
// 	maxVal := 0
// 	for _, val := range dp {
// 		maxVal = max(val, maxVal)
// 	}
// 	return maxVal + 1
// }

func findLongestChain(pairs [][]int) int {
	sort.Slice(pairs, func(i, j int) bool {
		return pairs[i][1] < pairs[j][1]
	})
	last := -2000
	length := 0
	for _, val := range pairs {
		if val[0] > last {
			length += 1
			last = val[1]
		}
	}
	return length
}

func TestChain(t *testing.T) {
	assert := require.New(t)
	assert.Equal(2, findLongestChain([][]int{
		{1, 2}, {2, 3}, {3, 4},
	}))
	assert.Equal(3, findLongestChain([][]int{
		{1, 2}, {7, 8}, {4, 5},
	}))
	assert.Equal(4, findLongestChain([][]int{
		{1, 2}, {3, 7}, {4, 5}, {6, 7}, {8, 10},
	}))
}
