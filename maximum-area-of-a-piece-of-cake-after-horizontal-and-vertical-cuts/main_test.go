package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 10e8 + 7

func maxSlice(cuts []int) int {
	prev, max := 0, 0
	for _, cut := range cuts {
		tmp := cut - prev
		if max < tmp {
			max = tmp
		}
		prev = cut
	}
	return max
}

func maxArea(h int, w int, horizontalCuts []int, verticalCuts []int) int {
	sort.Ints(horizontalCuts)
	sort.Ints(verticalCuts)
	horizontalCuts = append(horizontalCuts, h)
	verticalCuts = append(verticalCuts, w)
	return (maxSlice(horizontalCuts) * maxSlice(verticalCuts)) % mod
}

func TestMaxArea(t *testing.T) {
	assert := require.New(t)
	assert.True(true)
	assert.Equal(4, maxArea(
		5, 4, []int{1, 2, 4}, []int{1, 3},
	))
}
