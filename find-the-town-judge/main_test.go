package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func findJudge(N int, trust [][]int) int {
	stateList := make([]int, N)

	for _, trustState := range trust {
		stateList[trustState[0]-1] = -N
		if stateList[trustState[1]-1] != -1 {
			stateList[trustState[1]-1] += 1
		}
	}

	candidate := -1
	for idx, state := range stateList {
		if state == N-1 {
			if candidate != -1 {
				return -1
			}
			candidate = idx + 1
		}
	}
	return candidate
}

func TestFindJudge(t *testing.T) {
	assert.Equal(t, 2, findJudge(2, [][]int{{1, 2}}), nil)
	assert.Equal(t, 3, findJudge(3, [][]int{{1, 3}, {2, 3}}), nil)
	assert.Equal(t, -1, findJudge(3, [][]int{{1, 3}, {2, 3}, {3, 1}}), nil)
	assert.Equal(t, -1, findJudge(3, [][]int{{1, 2}, {2, 3}}), nil)
	assert.Equal(t, 3, findJudge(4, [][]int{{1, 3}, {1, 4}, {2, 3}, {2, 4}, {4, 3}}), nil)
	assert.Equal(t, -1, findJudge(2, [][]int{{1, 2}, {2, 1}, {1, 2}}), nil)
}
