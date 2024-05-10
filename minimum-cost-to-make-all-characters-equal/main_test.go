package main

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
)

func min(a, b int64) int64 {
	if a < b {
		return a
	}
	return b
}

func minimumCost(s string) int64 {
	dpFront := make([][2]int64, len(s))
	prev := [2]int64{0, 0}
	for i := 0; i < len(dpFront); i++ {
		n := int(s[i] - '0')
		dpFront[i][n] = min(
			prev[n],
			prev[n^1]+int64(i),
		)
		dpFront[i][n^1] = min(
			prev[n]+int64(i+1),
			prev[n^1]+int64(i*2+1),
		)

		prev = dpFront[i]
	}
	// fmt.Printf("dpFront: %v\n", dpFront)
	dpEnd := make([][2]int64, len(s))
	prev = [2]int64{0, 0}
	for i := len(dpEnd) - 1; i >= 0; i-- {
		n := int(s[i] - '0')
		dpEnd[i][n] = min(
			prev[n],
			prev[n^1]+int64(len(dpEnd)-i-1),
		)
		dpEnd[i][n^1] = min(
			prev[n]+int64(len(dpEnd)-i),
			prev[n^1]+int64((len(dpEnd)-i)*2+1),
		)

		prev = dpEnd[i]
	}
	minVal := int64(math.MaxInt64)
	for i := range dpFront {
		minVal = min(
			minVal,
			min(dpFront[i][0]+dpEnd[i][0], dpFront[i][1]+dpEnd[i][1]),
		)
	}
	// fmt.Printf("dpEnd: %v\n", dpEnd)
	return minVal
}

func TestCost(t *testing.T) {
	assert := require.New(t)
	assert.Equal(int64(9), minimumCost("010101"))
	assert.Equal(int64(2), minimumCost("0011"))
}
