package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e+9 + 7

func IsPrime(b byte) bool {
	return b == '2' || b == '3' || b == '5' || b == '7'
}

func beautifulPartitions(s string, k int, minLength int) int {
	if !IsPrime(s[0]) || IsPrime(s[len(s)-1]) {
		return 0
	}
	if k == 1 {
		return 1
	}
	dp := make([]int, 0)
	cutPoints := []int{}
	prev := 0
	for i := 1; i < len(s); i++ {
		if IsPrime(s[i]) && !IsPrime(s[i-1]) &&
			len(s)-i >= minLength && i >= minLength {
			cutPoints = append(cutPoints, i)
			dp = append(dp, prev+1)
			prev++
		}
	}
	if len(dp) == 0 {
		return 0
	}
	// fmt.Printf("dp: %v\n", dp)
	k--
	next := make([]int, len(cutPoints))
	for ; k != 1; k-- {
		prev := 0
		for idx, point := range cutPoints {
			next[idx] = 0
			for i := idx - 1; i >= 0; i-- {
				if point-cutPoints[i] >= minLength {
					// fmt.Printf("point: %v, cutPoints[i]: %v, i: %v\n", point, cutPoints[i], i)
					next[idx] += dp[i]
					break
				}
			}
			next[idx] += prev
			next[idx] %= mod
			prev = next[idx]
		}
		dp, next = next, dp
		// fmt.Printf("dp: %v\n", dp)
	}
	return dp[len(dp)-1]
}

func TestBe(t *testing.T) {
	assert := require.New(t)
	assert.Equal(1, beautifulPartitions("23542185131", 3, 3))
	assert.Equal(3, beautifulPartitions("23542185131", 3, 2))
	assert.Equal(3, beautifulPartitions("23542185131", 2, 2))
	assert.Equal(1, beautifulPartitions("23542185131", 1, 2))
}
