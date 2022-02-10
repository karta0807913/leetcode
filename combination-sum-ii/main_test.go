package main

import (
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func recursive(candidates, history []int, target int) (ans [][]int) {
	// fmt.Printf("history: %v, target: %v\n", history, target)
	ans = [][]int{}
	for i := range candidates {
		if candidates[i] > target {
			return
		} else if i != 0 && candidates[i] == candidates[i-1] {
			continue
		} else if candidates[i] == target {
			ans = append(ans, append([]int{}, append(history, candidates[i])...))
		}
		ans = append(ans,
			recursive(candidates[i+1:], append(history, candidates[i]), target-candidates[i])...,
		)
	}
	return
}

func combinationSum2(candidates []int, target int) (ans [][]int) {
	sort.Ints(candidates)
	// fmt.Printf("candidates: %v\n", candidates)
	return recursive(candidates, make([]int, 0, len(candidates)), target)
}

func TestCombine(t *testing.T) {
	assert.Equal(t, [][]int{
		{1, 1, 6}, {1, 2, 5}, {1, 7}, {2, 6},
	}, combinationSum2([]int{10, 1, 2, 7, 6, 1, 5}, 8))
	assert.Equal(t, [][]int{
		{1, 2, 3}, {2, 4},
	}, combinationSum2([]int{1, 2, 3, 4}, 6))
	assert.Equal(t, [][]int{
		{12},
	}, combinationSum2([]int{12, 12, 13, 14}, 12))
	assert.Equal(t, [][]int{
		{12, 12},
	}, combinationSum2([]int{12, 12, 13, 14}, 24))
	assert.Equal(t, [][]int{
		{12, 13},
	}, combinationSum2([]int{12, 12, 13, 14}, 25))
	assert.Equal(t, [][]int{
		{1, 1},
	}, combinationSum2([]int{1, 1}, 2))
	assert.Equal(t, [][]int{
		{12},
	}, combinationSum2([]int{12, 12}, 12))
	assert.Equal(t, [][]int{
		{12},
	}, combinationSum2([]int{12}, 12))
	assert.Equal(t, [][]int{}, combinationSum2([]int{13}, 12))
	assert.Equal(t, [][]int{
		{1, 3}, {4},
	}, combinationSum2([]int{5, 1, 5, 4, 3, 3}, 4))
}
