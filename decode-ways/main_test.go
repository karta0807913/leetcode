package main

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func numDecodingsRecursive(s string) int {
	if len(s) == 0 {
		return 1
	}
	if s[0] == '0' {
		return 0
	}
	if len(s) == 1 {
		return 1
	}
	if s[0] == '1' {
		return numDecodingsRecursive(s[1:]) + numDecodingsRecursive(s[2:])
	}
	if s[0] == '2' {
		if '0' <= s[1] && s[1] <= '6' {
			return numDecodingsRecursive(s[1:]) + numDecodingsRecursive(s[2:])
		}
	}
	return numDecodingsRecursive(s[1:])
}

func numDecodingsOld(s string) int {
	if s[0] == '0' {
		return 0
	}
	pprev, prev := 1, 1
	for idx := 1; idx < len(s); idx++ {
		current := 0
		// one digit
		if s[idx] != '0' {
			current += prev
		}
		// tow digits
		if s[idx-1] == '1' || (s[idx-1] == '2' && '0' <= s[idx] && s[idx] <= '6') {
			current += pprev
		}
		if current == 0 {
			return 0
		}
		pprev, prev = prev, current
	}
	return prev
}

// more readable
func numDecodings(s string) int {
	if s[0] == '0' {
		return 0
	}
	canEncode := make(map[string]bool, 26)
	for i := 0; i < 26; i++ {
		canEncode[fmt.Sprintf("%d", i+1)] = true
	}
	prev, current := 1, 1
	for i := 1; i < len(s); i++ {
		next := 0
		if canEncode[s[i:i+1]] {
			next += current
		}
		if canEncode[s[i-1:i+1]] {
			next += prev
		}
		prev, current = current, next
	}
	return current
}

func TestDecodings(t *testing.T) {
	assert.Equal(t, 2, numDecodings("12"))
	assert.Equal(t, 3, numDecodings("226"))
	assert.Equal(t, 0, numDecodings("06"))
	assert.Equal(t, 2, numDecodings("11106"))
	assert.Equal(t, 2, numDecodings("11102"))
	assert.Equal(t, 1, numDecodings("27"))
	assert.Equal(t, 9, numDecodings("123123"))
	assert.Equal(t, 3, numDecodings("1231"))
	assert.Equal(t, 6, numDecodings("12312"))
	assert.Equal(t, 0, numDecodings("10011"))
	assert.Equal(t, 0, numDecodings("230"))
}
