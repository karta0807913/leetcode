package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func superEggDrop(k int, n int) int {
	if k == 1 {
		return n
	}
	dp := make([]int, 50)
	for i := range dp {
		dp[i] = i
	}
	for ; k > 1; k-- {
		dpNext := make([]int, 1, len(dp))
		dpNext[0] = 0
		for m := 1; dpNext[m-1] < n; m++ {
			dpNext = append(dpNext, dp[m-1]+dpNext[m-1]+1)
		}
		dp = dpNext
		// fmt.Println(len(dp))
	}
	// fmt.Printf("dp: %v\n", dp)
	return len(dp) - 1
}

func TestEgg(t *testing.T) {
	assert.Equal(t, 3, superEggDrop(2, 6))
	assert.Equal(t, 4+1, superEggDrop(3, 14))
}
