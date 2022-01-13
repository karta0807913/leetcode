package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

type IntSlice [][]int

func (s IntSlice) Len() int           { return len(s) }
func (s IntSlice) Less(i, j int) bool { return s[i][0] < s[j][0] }
func (s IntSlice) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }

func merge(intervals [][]int) (ans [][]int) {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] < intervals[j][0]
	})
	// sort.Sort(IntSlice(intervals))
	current := intervals[0]
	for idx := 1; idx < len(intervals); idx++ {
		if intervals[idx][1] <= current[1] {
		} else if intervals[idx][0] <= current[1] {
			current[1] = intervals[idx][1]
		} else {
			ans = append(ans, current)
			current = intervals[idx]
		}
	}
	ans = append(ans, current)
	return
}

func TestMerge(t *testing.T) {
	assert.ElementsMatch(t, [][]int{
		{1, 6}, {8, 10}, {15, 18},
	}, merge([][]int{
		{1, 3}, {2, 6}, {8, 10}, {15, 18},
	}))
	assert.ElementsMatch(t, [][]int{
		{1, 6},
	}, merge([][]int{
		{1, 6},
	}))
	assert.ElementsMatch(t, [][]int{
		{1, 6},
	}, merge([][]int{
		{1, 6}, {2, 3},
	}))
	assert.ElementsMatch(t, [][]int{
		{1, 7},
	}, merge([][]int{
		{1, 6}, {6, 7},
	}))
	assert.ElementsMatch(t, [][]int{
		{1, 7},
	}, merge([][]int{
		{1, 6}, {1, 7},
	}))
	assert.ElementsMatch(t, [][]int{
		{1, 6}, {7, 8},
	}, merge([][]int{
		{1, 6}, {7, 8},
	}))
}
