package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func restoreArray(adjacentPairs [][]int) (ans []int) {
	relatedMap := make(map[int][]int)

	for _, val := range adjacentPairs {
		relatedMap[val[0]] = append(relatedMap[val[0]], val[1])
		relatedMap[val[1]] = append(relatedMap[val[1]], val[0])
	}
	nextPoint := 0
	for key, val := range relatedMap {
		if len(val) == 1 {
			nextPoint = key
			break
		}
	}
	ans = append(ans, nextPoint)
	nextPoint = relatedMap[nextPoint][0]
	for {
		ans = append(ans, nextPoint)
		between := relatedMap[nextPoint]
		if len(between) == 1 {
			break
		} else if between[0] == ans[len(ans)-2] {
			nextPoint = between[1]
		} else {
			nextPoint = between[0]
		}
	}
	return
}

func TestRestore(t *testing.T) {
	assert.Equal(t, []int{1, 2, 3, 4}, restoreArray([][]int{
		{2, 1}, {3, 4}, {3, 2},
	}))
}
