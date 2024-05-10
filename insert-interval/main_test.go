package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func insert(intervals [][]int, newInterval []int) [][]int {
	result := make([][]int, 0, len(intervals)+1)
	i := 0
	for ; i < len(intervals); i++ {
		if intervals[i][1] >= newInterval[0] {
			break
		}
		result = append(result, intervals[i])
	}

	for i != len(intervals) &&
		(intervals[i][0] <= newInterval[1]) {
		if intervals[i][0] < newInterval[0] {
			newInterval[0] = intervals[i][0]
		}
		if newInterval[1] < intervals[i][1] {
			newInterval[1] = intervals[i][1]
		}
		i++
	}
	result = append(result, newInterval)

	for ; i < len(intervals); i++ {
		result = append(result, intervals[i])
	}

	return result
}

func TestInsert(t *testing.T) {
	assert := require.New(t)
	assert.Equal(
		[][]int{{1, 5}, {6, 8}},
		insert([][]int{{1, 5}}, []int{6, 8}),
	)
	assert.Equal(
		[][]int{{1, 5}},
		insert([][]int{{1, 5}}, []int{2, 3}),
	)
	assert.Equal(
		[][]int{{1, 5}, {6, 9}},
		insert([][]int{{1, 3}, {6, 9}}, []int{2, 5}),
	)
	assert.Equal(
		[][]int{
			{1, 2}, {3, 10}, {12, 16},
		},
		insert([][]int{
			{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16},
		}, []int{4, 8}),
	)
}
