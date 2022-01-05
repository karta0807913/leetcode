package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/assert"
)

func max(a, b int) int {
	if a < b {
		return b
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func calc(a, b int, dp [][]int) int {
	if a >= b {
		return 0
	}
	if dp[a][b] != 0 {
		return dp[a][b]
	}
	dp[a][b] = math.MaxInt32
	for splitSpot := (a+b)/2 - 1; splitSpot < b; splitSpot++ {
		temp := max(calc(a, splitSpot-1, dp), calc(splitSpot+1, b, dp)) + splitSpot
		if dp[a][b] > temp {
			dp[a][b] = temp
		}
	}
	return dp[a][b]
}

func getMoneyAmount(n int) int {
	dp := make([][]int, n+1)
	for idx := range dp {
		dp[idx] = make([]int, n+1)
	}
	return calc(1, n, dp)
}

// type Operation struct {
// 	Cost           int
// 	OperationCount int
// }
// func getMoneyAmount(n int) int {
// 	dp := make([]Operation, max(n+1, 4))
// 	dp[1] = Operation{
// 		Cost:           0,
// 		OperationCount: 0,
// 	}
// 	dp[2] = Operation{
// 		Cost:           1,
// 		OperationCount: 1,
// 	}
// 	for num := 3; num <= n; num++ {
// 		dp[num].Cost = math.MaxInt64
// 		for left := 2; left < num; left++ {
// 			operation := Operation{
// 				Cost:           max(dp[num-left].Cost+dp[num-left].OperationCount*left, dp[left-1].Cost) + left,
// 				OperationCount: max(dp[num-left].OperationCount, dp[left-1].OperationCount) + 1,
// 			}
// 			if operation.Cost < dp[num].Cost {
// 				dp[num] = operation
// 			}
// 		}
// 	}
// 	return dp[n].Cost
// }

func TestMonyAmount(t *testing.T) {
	assert.Equal(t, 1, getMoneyAmount(2))
	assert.Equal(t, 2, getMoneyAmount(3))
	assert.Equal(t, 4, getMoneyAmount(4))
	assert.Equal(t, 6, getMoneyAmount(5))
	assert.Equal(t, 16, getMoneyAmount(10))
	assert.Equal(t, 18, getMoneyAmount(11))
	assert.Equal(t, 49, getMoneyAmount(20))
	assert.Equal(t, 400, getMoneyAmount(100))
	assert.Equal(t, 952, getMoneyAmount(200))
}
