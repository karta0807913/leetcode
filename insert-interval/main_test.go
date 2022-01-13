package main

import (
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

func isCovered(left []int, right []int) bool {
	return right[0] <= left[1]
}

func insert(intervals [][]int, newInterval []int) [][]int {
	if len(intervals) == 0 {
		return [][]int{newInterval}
	}
	left, right := 0, len(intervals)-1
	idx := 0
	for left <= right {
		i := (left + right) / 2
		if intervals[i][0] < newInterval[0] {
			left = i + 1
			idx = i
		} else if intervals[i][0] > newInterval[0] {
			right = i - 1
			idx = right
		} else {
			idx = i
			break
		}
	}

	var before [][]int
	if idx != -1 && isCovered(intervals[idx], newInterval) {
		newInterval[0] = min(intervals[idx][0], newInterval[0])
		idx -= 1
	}
	idx++
	before = intervals[:idx]
	for idx != len(intervals) && isCovered(newInterval, intervals[idx]) {
		newInterval[1] = max(intervals[idx][1], newInterval[1])
		idx++
	}

	return append(before, append([][]int{newInterval}, intervals[idx:]...)...)

}

func TestInsert(t *testing.T) {
	assert.Equal(t, [][]int{
		{0, 5}, {7, 16},
	}, insert([][]int{
		{0, 5}, {9, 12},
	}, []int{7, 16}))
	assert.Equal(t, [][]int{
		{3, 5}, {6, 8}, {9, 11},
	}, insert([][]int{
		{3, 5}, {6, 8},
	}, []int{9, 11}))
	assert.Equal(t, [][]int{
		{0, 2}, {3, 5}, {6, 8},
	}, insert([][]int{
		{3, 5}, {6, 8},
	}, []int{0, 2}))
	assert.Equal(t, [][]int{
		{0, 9},
	}, insert([][]int{
		{1, 5}, {6, 8},
	}, []int{0, 9}))
	assert.Equal(t, [][]int{
		{0, 5},
	}, insert([][]int{
		{1, 5},
	}, []int{0, 3}))
	assert.Equal(t, [][]int{
		{0, 5},
	}, insert([][]int{
		{0, 3},
	}, []int{3, 5}))
	assert.Equal(t, [][]int{
		{1, 6}, {7, 8},
	}, insert([][]int{
		{1, 2}, {3, 4}, {5, 6}, {7, 8},
	}, []int{2, 5}))
	assert.Equal(t, [][]int{
		{1, 9},
	}, insert([][]int{
		{1, 2}, {5, 9},
	}, []int{2, 5}))
	assert.Equal(t, [][]int{
		{1, 2}, {3, 4}, {5, 6},
	}, insert([][]int{
		{1, 2}, {5, 6},
	}, []int{3, 4}))
	assert.Equal(t, [][]int{
		{1, 3}, {5, 9},
	}, insert([][]int{
		{1, 2}, {5, 9},
	}, []int{2, 3}))
	assert.Equal(t, [][]int{
		{1, 3}, {5, 9},
	}, insert([][]int{
		{1, 3}, {5, 6},
	}, []int{6, 9}))
	assert.Equal(t, [][]int{
		{1, 2}, {3, 9},
	}, insert([][]int{
		{1, 2}, {3, 4},
	}, []int{4, 9}))
	assert.Equal(t, [][]int{
		{1, 2}, {3, 9}, {10, 11},
	}, insert([][]int{
		{1, 2}, {3, 4}, {10, 11},
	}, []int{4, 9}))
}
