package main

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func integerReplacement(n int) int {
	i := 0
	for n != 1 {
		if n == 3 {
			return i + 2
		}
		if n&1 == 1 {
			if n&2 == 2 {
				n += 1
			} else {
				n -= 1
			}
		} else {
			n >>= 1
		}
		// fmt.Printf("n: %v\n", n)
		i += 1
	}
	return i
}

func TestInteger(t *testing.T) {
	assert := require.New(t)
	assert.Equal(2, integerReplacement(3))
}
