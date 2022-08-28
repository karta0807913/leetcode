package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func brokenCalc(startValue int, target int) int {
	operations := 0
	for startValue < target {
		startValue <<= 1
		operations++
	}

	gap := startValue - target
	operations, gap = operations+(gap>>operations), gap&(^(-1 << operations))
	for gap != 0 {
		operations += gap & 1
		gap >>= 1
	}

	return operations
}

func TestCalc(t *testing.T) {
	assert := require.New(t)

	assert.Equal(6, brokenCalc(7, 15))
	assert.Equal(2, brokenCalc(2, 3))
	assert.Equal(2, brokenCalc(5, 8))
	assert.Equal(3, brokenCalc(3, 10))
	assert.Equal(4, brokenCalc(4, 10))
	assert.Equal(458697, brokenCalc(1919127, 373872631))
}
