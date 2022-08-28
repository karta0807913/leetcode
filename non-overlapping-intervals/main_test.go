package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/require"
)

func eraseOverlapIntervals(intervals [][]int) int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][1] < intervals[j][1]
	})
	end := intervals[0]
	count := 1
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] >= end[1] {
			end = intervals[i]
			count += 1
		}
	}
	return len(intervals) - count
}

func TestOverlap(t *testing.T) {
	assert := require.New(t)
	assert.Equal(1, eraseOverlapIntervals([][]int{
		{1, 2}, {2, 3}, {3, 4}, {1, 3},
	}))
	assert.Equal(1, eraseOverlapIntervals([][]int{
		{1, 4},
		{3, 5},
		{4, 7},
	}))
}
