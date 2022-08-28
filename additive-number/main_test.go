package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func recursive(prev, want int, num string) bool {
	// fmt.Printf("num: %v\n", num)
	if len(num) == 0 {
		return true
	}
	// fmt.Printf("want: %v\n", want)
	next := 0
	for idx := 1; idx <= len(num); idx++ {
		next *= 10
		next += int(num[idx-1] - '0')
		// fmt.Printf("next: %v\n", next)
		if next < want {
			if next == 0 {
				return false
			}
			continue
		} else if next > want {
			return false
		}
		if recursive(next, next+prev, num[idx:]) {
			return true
		}
	}

	return false
}

func isAdditiveNumber(num string) bool {
	var prev = int(num[0] - '0')
	for i := 1; i < len(num); i++ {
		var current = int(num[i] - '0')
		for j := i + 1; j < len(num); j++ {
			// fmt.Printf("prev: %v, current: %v\n", prev, current)
			if recursive(current, prev+current, num[j:]) {
				return true
			}
			if current == 0 {
				break
			}
			current = current*10 + int(num[j]-'0')
		}
		if prev == 0 {
			return false
		}
		prev = prev*10 + int(num[i]-'0')
	}
	return false
}

func TestNumber(t *testing.T) {
	assert := require.New(t)
	assert.False(isAdditiveNumber("0235813"))
	assert.False(isAdditiveNumber("1203"))
	assert.True(isAdditiveNumber("000"))
	assert.True(isAdditiveNumber("101"))
	assert.True(isAdditiveNumber("11011"))
	assert.False(isAdditiveNumber("110112"))
	assert.False(isAdditiveNumber("12"))
	assert.True(isAdditiveNumber("123"))
	assert.True(isAdditiveNumber("199100199"))
	assert.True(isAdditiveNumber("112358"))
}
