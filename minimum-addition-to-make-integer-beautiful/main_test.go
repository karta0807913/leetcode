package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func recursive(n, digit int64, sum, target int) (bool, int64) {
	if digit == 0 {
		return true, 0
	}
	tmp := n / digit
	sum += int(tmp)
	if sum > target {
		return false, 0
	}
	ok, remains := recursive(n%digit, digit/10, sum, target)
	if !ok {
		if tmp == 9 {
			return false, 0
		}
		tmp++
		sum++
		if sum > target {
			return false, 0
		}
		return true, digit*tmp + remains
	}
	return true, digit*tmp + remains
}

func makeIntegerBeautiful(n int64, target int) int64 {
	digit := int64(1)
	for n/digit != 0 {
		digit *= 10
	}
	digit /= 10
	ok, remains := recursive(n, digit, 0, target)
	if ok {
		return remains - n
	}
	return digit*10 - n
}

func TestBeautiful(t *testing.T) {
	assert := require.New(t)
	assert.Equal(int64(2), makeIntegerBeautiful(8, 2))
}
