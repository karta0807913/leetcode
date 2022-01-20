package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func canCompleteCircuit(gas []int, cost []int) int {
	gasTank := 0
	startStation := 0
	total := 0
	for idx := range gas {
		cost := gas[idx] - cost[idx]
		gasTank += cost
		total += cost
		if gasTank < 0 {
			startStation = idx + 1
			gasTank = 0
		}
	}
	if total >= 0 {
		return startStation
	}
	return -1
}

func TestCircult(t *testing.T) {
	assert.Equal(t, -1, canCompleteCircuit([]int{
		3, 2, 4,
	}, []int{
		3, 4, 3,
	}))
	assert.Equal(t, 0, canCompleteCircuit([]int{
		1, 1, 1,
	}, []int{
		1, 1, 1,
	}))
	assert.Equal(t, 3, canCompleteCircuit([]int{
		1, 2, 3, 4, 5,
	}, []int{
		3, 4, 5, 1, 2,
	}))
	assert.Equal(t, -1, canCompleteCircuit([]int{
		4,
	}, []int{
		5,
	}))
}
