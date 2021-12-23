package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func uniquePaths(m int, n int) int {
	m, n = m-1, n-1
	if m < n {
		m, n = n, m
	}
	result := 1
	for i := m + 1; i <= m+n; i++ {
		result *= i
	}
	devide := 1
	for i := 1; i <= n; i++ {
		devide *= i
	}
	return result / devide
}

func TestPaths(t *testing.T) {
	assert.Equal(t, 3, uniquePaths(3, 2))
	assert.Equal(t, 28, uniquePaths(3, 7))
}
