package main

import "math"

const (
	Holding    = 0
	NotHolding = 1
)

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

func maximumStrength(nums []int, k int) int64 {
	// [k][index][2]{HoldingMax, NotHoldingMax}
	dp := make([][][2]int64, k+1)
	for i := range dp {
		dp[i] = make([][2]int64, len(nums))
		for j := range dp[i] {
			dp[i][j] = [2]int64{math.MinInt64, math.MinInt64}
		}
	}

	previousN := [2]int64{0, 0}
	for idx := range nums {
		previousK := [2]int64{0, 0}
		for kCount := 0; kCount < k && kCount < idx; k++ {
			// previous add
			dp[kCount][idx][NotHolding] = max(previousN[NotHolding], previousK[Holding])
			// add element or new holding
			dp[kCount][idx][Holding] = max(previousK[Holding] + nums[idx])
			// new holding
			previousK = dp[kCount][idx]
		}
	}
}
