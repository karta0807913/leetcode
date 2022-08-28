package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

const mod = 1e9 + 7

func idealArrays(n int, maxValue int) int {
}

func TestArrays(t *testing.T) {
	assert := require.New(t)
	assert.Equal(10, idealArrays(2, 5))
	assert.Equal(11, idealArrays(5, 3))
}
