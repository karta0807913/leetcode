package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func countNumbersWithUniqueDigits(n int) int {
	current := 1
	for idx := 0; idx < n; idx++ {
		t := 1
		for num := 0; num < idx; num++ {
			t *= 9 - num
		}
		t *= 9
		current = current + t
	}
	return current
}

func TestDigits(t *testing.T) {
	assert.Equal(t, 1, countNumbersWithUniqueDigits(0))
	assert.Equal(t, 10, countNumbersWithUniqueDigits(1))
	assert.Equal(t, 91, countNumbersWithUniqueDigits(2))
	assert.Equal(t, 9*9*8+1*9*9+1*1*9+1, countNumbersWithUniqueDigits(3))
	assert.Equal(t, 9*9*8*7+1*9*9*8+1*1*9*9+1*1*1*9+1, countNumbersWithUniqueDigits(4))
}

// 9+1, n = 1
// 9*9+1*9+1, n = 2
// 9*9*8+1*9*9+1*1*9+1, n = 3
// 9*9*8*7+1*9*9*8+1*1*9*9+1*1*1*9+1, n = 4
